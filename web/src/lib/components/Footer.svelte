<script lang="ts">
    import { m } from '$lib/paraglide/messages.js';
    import config from '$lib/config';
</script>

<footer class="flex items-center justify-between">
    <p>{config.appName}ㆍ{config.appVersion}</p>
    <p class="text-right">{ m['dashboard.footer.copyright']()}<br> { m['dashboard.footer.author']() }</p>
</footer>