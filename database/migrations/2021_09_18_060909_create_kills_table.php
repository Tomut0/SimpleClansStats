<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateKillsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('sc_kills', function (Blueprint $table) {
            $table->bigIncrements('kill_id');
            $table->string('attacker', 40);
            $table->string('attacker_tag', 25);
            $table->string('victim', 40);
            $table->string('victim_tag', 25);
            $table->char('kill_type', 1);
            $table->string('attacker_uuid', 38);
            $table->string('victim_uuid', 38);
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('sc_kills');
    }
}
