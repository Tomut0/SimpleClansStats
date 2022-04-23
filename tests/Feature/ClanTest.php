<?php

namespace Tests\Feature;

use App\Clan;
use Tests\TestCase;
use Throwable;

class ClanTest extends TestCase
{
    private $clan;

    public function setUp(): void
    {
        parent::setUp();
        $this->clan = factory(Clan::class)->make();
    }

    /**
     * @throws Throwable
     */
    public function tearDown(): void
    {
        parent::tearDown();
        $this->clan->delete();
    }

    /**
     * Tests if @link Clan is not null
     */
    public function testClanNonNull()
    {
        self::assertNotNull($this->clan);
    }
}
