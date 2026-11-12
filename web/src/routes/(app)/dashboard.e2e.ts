import { expect, test } from '@playwright/test';

test('edits the widget catalog and saves the dashboard layout', async ({ page }) => {
	let savedLayout:
		| { version: number; widgets: { id: string; x: number; y: number; w: number; h: number }[] }
		| undefined;
	await page.route('**/*', async (route) => {
		const url = new URL(route.request().url());
		if (url.pathname.endsWith('/auth/me')) {
			return route.fulfill({
				status: 200,
				contentType: 'application/json',
				body: JSON.stringify({ email: 'test@example.com', username: 'test', role: 'admin' })
			});
		}
		if (url.pathname.endsWith('/servers')) {
			return route.fulfill({
				status: 200,
				contentType: 'application/json',
				body: JSON.stringify([
					{
						id: '6ecad9ed-b8b1-49f2-8443-3b73203ba033',
						displayName: 'Test server',
						disabled: false,
						online: true,
						lastIngestAt: '2026-11-11T00:00:00Z'
					}
				])
			});
		}
		if (url.pathname.endsWith('/settings')) {
			return route.fulfill({
				status: 200,
				contentType: 'application/json',
				body: JSON.stringify({
					appName: 'SCStats',
					removeAuthor: false,
					removeTelemetry: false,
					timeZone: 'Europe/Moscow',
					dateFormat: 'DD.MM.YYYY',
					timeFormat: '24h',
					logoVersion: 0,
					serverOnlineTimeoutMinutes: 10,
					ingestRetentionDays: 30
				})
			});
		}
		if (url.pathname.endsWith('/dashboard/layout') && route.request().method() === 'GET') {
			return route.fulfill({
				status: 200,
				contentType: 'application/json',
				body: JSON.stringify(savedLayout ?? { version: 1, widgets: [] })
			});
		}
		if (url.pathname.endsWith('/dashboard/layout') && route.request().method() === 'PUT') {
			savedLayout = JSON.parse(route.request().postData() ?? '{}');
			return route.fulfill({ status: 204, body: '' });
		}
		if (url.pathname.endsWith('/dashboard/events')) {
			return route.fulfill({
				status: 200,
				contentType: 'text/event-stream',
				body:
					'event: snapshot\ndata: {"state":{"clans":[{"tag":"RME","balance":752151,"status":"verified"}],"players":[{"name":"Minat0"}],"extensions":{"scCollectorV1":{"wars":[],"onlinePlayers":[],"killsByType":[{"type":"allies","count":3},{"type":"rivals","count":2}]}}},"events":[]}\n\n' +
					'event: snapshot\ndata: {"state":{"clans":[{"tag":"RME","balance":752151,"status":"verified"},{"tag":"ABC","balance":125522,"status":"unverified"}],"players":[{"name":"Minat0"},{"name":"RoinujNosde"}],"extensions":{"scCollectorV1":{"wars":[],"onlinePlayers":[],"killsByType":[{"type":"allies","count":3},{"type":"rivals","count":2}]}}},"events":[]}\n\n'
			});
		}
		return route.continue();
	});

	let documentLoads = 0;
	page.on('load', () => documentLoads++);
	await page.goto('/');
	await expect(page.locator('.grid-stack-item')).toHaveCount(9);
	const grid = page.locator('.grid-stack');
	await expect(grid).toHaveAttribute('gs-static', 'true');
	await expect(page.locator('[data-widget-id="clans"]')).toHaveClass(/ui-draggable-disabled/);
	await expect(page.locator('[data-widget-id="clans"]')).toHaveClass(/ui-resizable-disabled/);
	await expect(page.locator('[data-widget-id="clans"]')).toHaveAttribute('gs-h', '2');
	await expect(page.locator('[data-widget-id="clans"] strong')).toHaveText('2');
	await expect(page.locator('[data-widget-id="top"] table')).toBeVisible();
	await expect(page.getByText('Live updates connected')).toHaveCount(0);
	await page.screenshot({ path: 'test-results/dashboard.png' });
	documentLoads = 0;
	await page.waitForTimeout(11_000);
	expect(documentLoads).toBe(0);
	await page.getByRole('button', { name: 'Customize dashboard' }).click();
	await expect(grid).not.toHaveAttribute('gs-static', 'true');
	const clansWidget = page.locator('[data-widget-id="clans"]');
	await expect(clansWidget).not.toHaveClass(/ui-draggable-disabled/);
	await page.getByRole('button', { name: 'Remove Kills by type' }).click();
	const bounds = await grid.boundingBox();
	expect(bounds).not.toBeNull();
	await page.mouse.move(bounds!.x + bounds!.width * 0.5, bounds!.y + 3 * 76);
	await expect(page.locator('.add-widget-hover')).toHaveCount(0);
	await page.mouse.move(bounds!.x + bounds!.width * 0.12, bounds!.y + 3 * 76);
	const hoverTile = page.locator('.add-widget-hover');
	await expect(hoverTile).toBeVisible();
	const tileBounds = await hoverTile.boundingBox();
	expect(tileBounds).not.toBeNull();
	await page.mouse.click(
		tileBounds!.x + tileBounds!.width / 2,
		tileBounds!.y + tileBounds!.height / 2
	);
	await page.keyboard.press('Escape');
	const playersWidget = page.locator('[data-widget-id="players"]');
	const playersHeader = playersWidget.locator('.widget-header');
	await playersHeader.hover();
	const playersHeaderBounds = await playersHeader.boundingBox();
	const gridBounds = await grid.boundingBox();
	expect(playersHeaderBounds).not.toBeNull();
	expect(gridBounds).not.toBeNull();
	const beforePosition = [
		await playersWidget.getAttribute('gs-x'),
		await playersWidget.getAttribute('gs-y')
	];
	await page.mouse.move(
		playersHeaderBounds!.x + playersHeaderBounds!.width / 2,
		playersHeaderBounds!.y + playersHeaderBounds!.height / 2
	);
	await playersHeader.evaluate((element) =>
		element.addEventListener(
			'mousedown',
			(event) =>
				element.setAttribute('data-test-mousedown-prevented', String(event.defaultPrevented)),
			{ once: true }
		)
	);
	await page.mouse.down();
	await expect(playersHeader).toHaveAttribute('data-test-mousedown-prevented', 'true');
	await page.mouse.move(gridBounds!.x + gridBounds!.width / 12, gridBounds!.y + 3.5 * 76, {
		steps: 8
	});
	await page.mouse.up();
	await expect
		.poll(async () => [
			await playersWidget.getAttribute('gs-x'),
			await playersWidget.getAttribute('gs-y')
		])
		.not.toEqual(beforePosition);
	await page.getByRole('button', { name: 'Add widget' }).first().click();
	let addDialog = page.getByRole('dialog');
	await addDialog.getByRole('button', { name: 'Distribution' }).click();
	await addDialog.getByLabel('Field').selectOption('killTypes');
	await addDialog.getByRole('button', { name: 'Add widget' }).click();
	const distributionWidget = page.locator('[data-widget-id^="widget-"]');
	await expect(distributionWidget).toHaveAttribute('gs-w', '3');
	await expect(distributionWidget).toHaveAttribute('gs-h', '2');
	await page.getByRole('button', { name: 'Add widget' }).first().click();
	addDialog = page.getByRole('dialog');
	await addDialog.getByRole('button', { name: 'Counter' }).click();
	await addDialog.getByLabel('Field').selectOption('players');
	await addDialog.getByLabel('Format').selectOption('compact');
	await addDialog.getByRole('button', { name: 'Add widget' }).click();
	const customCounter = page.locator('[data-widget-id^="counter-"]');
	await expect(customCounter).toHaveCount(1);
	await expect(customCounter.locator('strong')).toHaveText('2');
	await clansWidget.hover();
	const resizeHandle = clansWidget.locator('.ui-resizable-se');
	await expect(resizeHandle).toBeVisible();
	const beforeWidth = Number(await clansWidget.getAttribute('gs-w'));
	const resizeBounds = await resizeHandle.boundingBox();
	expect(resizeBounds).not.toBeNull();
	await page.mouse.move(
		resizeBounds!.x + resizeBounds!.width / 2,
		resizeBounds!.y + resizeBounds!.height / 2
	);
	await page.mouse.down();
	await page.mouse.move(
		resizeBounds!.x + resizeBounds!.width / 2 + 100,
		resizeBounds!.y + resizeBounds!.height / 2 + 76,
		{ steps: 5 }
	);
	await page.mouse.up();
	await expect.poll(() => clansWidget.getAttribute('gs-w')).not.toBe(String(beforeWidth));
	await page.getByRole('button', { name: 'Save dashboard' }).click();
	await expect(grid).toHaveAttribute('gs-static', 'true');
	await expect(clansWidget).toHaveClass(/ui-resizable-disabled/);
	await expect.poll(() => savedLayout?.widgets.length).toBe(10);
	expect(savedLayout?.widgets.map((widget) => widget.id)).not.toContain('add-widget');
	expect(savedLayout?.widgets).toContainEqual(
		expect.objectContaining({ type: 'counter', field: 'players', format: 'compact' })
	);
	await page.reload();
	await expect(page.locator('.grid-stack-item')).toHaveCount(10);
	await expect(page.locator('[data-widget-id="clans"]')).toHaveAttribute(
		'gs-w',
		String(savedLayout!.widgets.find((widget) => widget.id === 'clans')!.w)
	);
});
