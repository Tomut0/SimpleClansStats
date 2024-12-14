<?php

namespace App\Console\Commands;

use App\Models\Clan;
use App\Models\ClanPlayer;
use Illuminate\Console\Command;
use Illuminate\Support\Facades\DB;

class UpdateDemoCommand extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'scstats:update-demo';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Updates SimpleClansStats demo data';

    /**
     * Execute the console command.
     */
    public function handle(): void
    {
        if (app()->environment() !== 'demo') {
            $this->comment("This command can only be run from demo mode");
            return;
        }

        Clan::all()->each(function (Clan $clan) {
            // update a clan's balance
            if (rand(0, 1)) {
                $this->comment("Updating {$clan->tag} balance...");
                $table = $clan->getTable();
                DB::connection('simpleclans')
                    ->update("update $table set balance = ? where tag = ?", [rand(0, 5000), $clan->tag]);
            }

            if (rand(0, 1)) {
                $clanPlayer = ClanPlayer::factory()->create([
                    'tag' => $clan->tag,
                ]);
                $this->comment("Adding {$clanPlayer->name} to {$clan->tag}...");
                $clanPlayer->clan()->associate($clan);
            } else {
                $clanPlayer = $clan->members()->get()->random()->first();
                $this->comment("Removing {$clanPlayer->name} from {$clan->tag}...");
                $clanPlayer->clan()->dissociate($clan);
            }
        });

        // create clans and clanplayers
        if (rand(0, 1) && Clan::all()->count() > 0) {
            $clan = Clan::all()->random(1)->first();
            DB::connection('simpleclans')->delete("delete from {$clan->getTable()} where tag = ?", [$clan->tag]);
            $this->comment("Deleting a clan $clan->tag");
        } else {
            $this->comment("Creating a clan");
            Clan::factory()->create()->each(function (Clan $clan) {
                ClanPlayer::factory()->count(rand(1, 3))->create([
                    'tag' => $clan->tag,
                ]);
            });
        }
    }
}
