@extends('layouts.html')
@section('title', env("APP_NAME"))
@section('content')
    <header>
        <div class="container-fluid">
            <nav class="navbar navbar-expand-lg navbar-light bg-light">
                <div class="collapse navbar-collapse">
                    <div class="container-fluid">
                        <ul class="navbar-nav navbar-nav me-auto mb-2 mb-lg-0">
                            <a class="navbar-brand" href="#">{{ env('APP_NAME') }}</a>

                            <li class="navbar-item"><a href="{{url('/')}}" class="nav-link active">{{ __('messages.header.home') }}</a></li>
                            <li class="navbar-item"><a href="{{url('/clans')}}" class="nav-link">{{ __('messages.header.clans') }}</a></li>
                            <li class="navbar-item"><a href="{{url('/players')}}" class="nav-link">{{ __('messages.header.players') }}</a></li>
                        </ul>
                    </div>
                </div>
            </nav>
        </div>
    </header>
    <main class="mt-3">
        <div class="container-fluid">
            <div class="row text-center">
                <div class="col">
                    <h4 class="bg-primary text-white p-3 mb-0">Top 10 Clans</h4>
                    <table class="table table-primary table-striped w-100">
                        <thead>
                            <tr>
                                <th scope="col">#</th>
                                <th scope="col">{{ __('messages.main.clanName') }}</th>
                                <th scope="col">{{ __('messages.main.clanMembers') }}</th>
                            </tr>
                        </thead>
                        <tbody>
{{--                        @foreach($clans->getTopClans() as $index => $clan)--}}
{{--                        <tr>--}}
{{--                            <th scope="row">{{ $clan->id }}</th>--}}
{{--                            <td>{{ $clan->name }}</td>--}}
{{--                            <td>{{ $clan->balance }}</td>--}}
{{--                        </tr>--}}
{{--                        @endforeach--}}
                        </tbody>
                    </table>
                </div>
                <div class="col">
                    <h4 class="bg-success text-white p-3 m-0">Top 10 Players</h4>
                    <table class="table table-success table-striped w-100">
                        <thead>
                        <tr>
                            <th scope="col">#</th>
                            <th scope="col">{{ __('messages.main.playerName') }}</th>
                            <th scope="col">{{ __('messages.main.playerKDR') }}</th>
                        </tr>
                        </thead>
                        <tbody>
                        @foreach($players->getTopPlayersByKDR() as $id => $player)
                        <tr>

                            <th scope="row">{{ $id+1 }}</th>
                            <td>{{ $player['Name'] }}</td>
                            <td>{{ $player['KDR'] }}</td>
                        </tr>
                        @endforeach
                        </tbody>
                    </table>
                </div>
                <div class="col">
                    <h4 class="bg-warning text-white p-3 m-0">Last 10 Kills</h4>
                    <table class="table table-warning table-striped w-100">
                        <thead>
                        <tr>
                            <th scope="col">#</th>
                            <th scope="col">{{ __('messages.main.attacker') }}</th>
                            <th scope="col">{{ __('messages.main.victim') }}</th>
                        </tr>
                        </thead>
                        <tbody>
                        @foreach($players->getLastPlayersKills() as $id => $player)
                        <tr>
                            <th scope="row">{{ $id+1 }}</th>
                            <td><span class="badge bg-dark .rounded-pill">ClanTag</span> {{ $player->attacker }}</td>
                            <td><span class="badge bg-dark .rounded-pill">ClanTag</span>{{ $player->victim }}</td>
                        </tr>
                        @endforeach
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </main>
    <footer class="text-muted py-5">
        <div class="container-fluid">
            <div class="float-end">
                <p class="mb-0">Made with ♥ by Minat0_</p>
                <p>Special for <a href="https://www.spigotmc.org/resources/simpleclans.71242/">SimpleClans</a></p>
            </div>
        </div>
    </footer>
@endsection
