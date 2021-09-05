<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateClansTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('sc_clans', function (Blueprint $table) {
            $table->bigIncrements('id');
            $table->tinyInteger('verified');
            $table->string('tag', 5);
            $table->string('name', 25);
            $table->text('description')->nullable();
            $table->tinyInteger('friendly_fire')->default(0);
            $table->bigInteger('founded');
            $table->bigInteger('last_used');
            $table->text('packed_allies')->nullable();
            $table->text('packed_rivals')->nullable();
            $table->mediumText('packed_bb');
            $table->text('cape_url')->nullable();
            $table->double('balance')->default(0.0);
            $table->tinyInteger('fee_enabled')->default(0);
            $table->double('fee_balance')->default(0.0);
            $table->text('ranks');
            $table->text('banner')->nullable();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('sc_clans');
    }
}
