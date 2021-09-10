<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreatePlayersTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('sc_players', function (Blueprint $table) {
            $table->bigIncrements('id');
            $table->string('tag', 16);
            $table->string('uuid', 36);
            $table->string('name', 16);
            $table->tinyInteger('leader')->default(0);
            $table->tinyInteger('friendly_fire')->default(0);
            $table->integer('neutral_kills')->nullable();
            $table->integer('rival_kills')->nullable();
            $table->integer('civilian_kills')->nullable();
            $table->integer('deaths')->nullable();
            $table->bigInteger('last_seen');
            $table->bigInteger('join_date');
            $table->tinyInteger('trusted')->default(0);
            $table->text('flags')->nullable();
            $table->text('packed_past_clans')->nullable();
            $table->text('resign_times')->nullable();
            $table->text('locale')->nullable();
            $table->text('ally_kills')->nullable();

        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('sc_players');
    }
}
