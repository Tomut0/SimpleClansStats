@extends('layouts.html')
@section('title', env("APP_NAME"))
@section('content')
    <main class="mt-3">
        <div class="container-fluid">
            <div class="row text-start">
                <div class="col">
                    {{--                    TODO: Translate table titles--}}
                    <h4 class="bg-primary text-white p-3 mb-0 text-center">Top 10 Clans</h4>
                    <table class="table table-primary table-striped w-100">
                        <thead>
                        <tr>
                            <th scope="col">#</th>
                            <th scope="col">{{ __('messages.main.clanName') }}</th>
                            <th scope="col">{{ __('messages.main.clanKDR')}}</th>
                            <th scope="col">{{ __('messages.main.clanMembers') }}</th>
                        </tr>
                        </thead>
                        <tbody>
                        @foreach($clans->getTopClans() as $index => $clan)
                            <tr>
                                <th scope="row">{{ $clan['rank'] }}</th>
                                <td class="modal-opener clan_name" data-tag='{{$clan["tag"]}}'>{{$clan['name']}}</td>
                                <td class="text-center">{{ number_format($clan['KDR'] ,2) }}</td>
                                <td class="text-center">{{ sizeof($clan['members']) }}</td>
                            </tr>
                        @endforeach
                        </tbody>
                    </table>
                </div>
                <div class="col">
                    <h4 class="bg-success text-white p-3 m-0 text-center">Top 10 Players</h4>
                    <table class="table table-success table-striped w-100">
                        <thead>
                        <tr>
                            <th scope="col">#</th>
                            <th scope="col">{{ __('messages.main.playerName') }}</th>
                            <th scope="col" class="text-center">{{ __('messages.main.playerKDR') }}</th>
                        </tr>
                        </thead>
                        <tbody>
                        @foreach($players->getTopPlayersByKDR() as $id => $player)
                            <tr>
                                <th scope="row">{{ $id+1 }}</th>
                                <td class="modal-opener" data-nick='{{$player["Name"]}}'>
                                    <img src="https://mc-heads.net/avatar/{{ $player['Name'] }}/16/"
                                         alt="{{ $player['Name'] }}'s Avatar"> {{ $player['Name'] }}</td>
                                <td class="text-center">{{ number_format($player['KDR'], 2) }}</td>
                            </tr>
                        @endforeach
                        </tbody>
                    </table>
                </div>
                <div class="col">
                    <h4 class="bg-warning text-white p-3 m-0 text-center">Last 10 Kills</h4>
                    <table class="table table-warning table-striped w-100">
                        <thead>
                        <tr>
                            <th scope="col">#</th>
                            <th scope="col">{{ __('messages.main.attacker') }}</th>
                            <th scope="col">{{ __('messages.main.victim') }}</th>
                        </tr>
                        </thead>
                        <tbody>
                        @foreach($players->getLastPlayersKills() as $id => $kill)
                            <tr>
                                <th scope="row">{{ $id+1 }}</th>
                                <td class="modal-opener player_name" data-nick='{{$kill->attacker}}'><span
                                        class="badge bg-dark .rounded-pill">{!! $kill->attacker_colored_tag !!}</span> {{ $kill->attacker }}
                                </td>
                                <td class="modal-opener player_name" data-nick='{{$kill->victim}}'><span
                                        class="badge bg-dark .rounded-pill">{!! $kill->victim_colored_tag !!}</span> {{ $kill->victim }}
                                </td>
                            </tr>
                        @endforeach
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </main>
@endsection
@push('scripts')
    <script src="{{asset("assets/js/detail_modal.js")}}" defer></script>
    <script defer src="{{asset("assets/js/trim_names.js")}}"></script>
@endpush
