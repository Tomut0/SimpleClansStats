@extends('layouts.html')
@section('title', env("APP_NAME"))
@section('content')
    <main>
        <div class="search-box block">
            <label>
                <span class="material-icons">search</span>
                <input type="text" placeholder="{{__('messages.main.search')}}">
            </label>
        </div>

        <div class="leaderboard block">
            <div class="heading">
                <h2>{{ __('messages.main.leaderboard') }}</h2>
                <div>
                    <h3>Clans</h3>
                    <h3>Players</h3>
                </div>
            </div>
            <nav class="sorted-by">
                <ul>
                    <li><span class="material-icons">account_balance</span></li>
                    <li><span class="material-icons">whatshot</span></li>
                    <li><span class="material-icons">people</span></li>
                </ul>
            </nav>
            <div class="board">
                <div class="cards">

                </div>
                <div class="list">

                </div>
            </div>
        </div>
        <div class="statistics block">

        </div>
    </main>
@endsection
