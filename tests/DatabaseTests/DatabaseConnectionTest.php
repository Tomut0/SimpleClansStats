<?php

namespace Tests\DatabaseTests;

use DB;
use Tests\TestCase;

class DatabaseConnectionTest extends TestCase
{
    /**
     * Tests if database connection is not null.
     *
     * @coversNothing
     */
    public function testConnection()
    {
        $connection = $this->getConnection();
        self::assertNotNull($connection);
    }

    /**
     * Tests if database contains app tables.
     *
     * @depends testConnection
     * @coversNothing
     */
    public function testTables()
    {
        $tables = ["sc_clans", "sc_players", "sc_kills"];

        foreach ($tables as $table) {
            $query = "SELECT * FROM $table";
            self::assertTrue(DB::statement($query), "Your database doesn't have $table table.");
        }
    }

    /**
     * Tests if mandatory tables have at least one row.
     *
     * @depends testTables
     * @coversNothing
     */
    public function testRowExists()
    {
        $tables = ["sc_clans", "sc_players"];

        foreach ($tables as $table) {
            self::assertTrue(DB::table($table)->exists());
        }
    }
}
