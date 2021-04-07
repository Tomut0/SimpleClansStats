@extends('layouts.modal')
@isset($player)
    @section('modal-title', $player['name'])
@section('modal-content')
    <table width="100%">
        <tbody>
        <tr>
            <td valign="top" width="150"><img src="https://mc-heads.net/player/{{$player['name']}}"
                                              alt="{{$player['name']}}'s skin"></td>
            <td>
                <table class="table table-striped">
                    <tr>
                        <th>Name</th>
                        <td>{{$player['name']}}</td>
                    </tr>
                    <tr>
                        <th>Join Date</th>
                        <td>{{ \App\Utils::formatDate($player['join_date']) }}</td>
                    </tr>
                    <tr>
                        <th>Last Seen</th>
                        <td>{{ \App\Utils::formatDateTime($player['last_seen']) }}</td>
                    </tr>
                    <tr>
                        <th>KDR</th>
                        <td>{{ \App\Utils::getFormattedKDR($player) }}</td>
                    </tr>
                    <tr>
                        <th>Rival Kills</th>
                        <td>{{ $player['rival_kills'] }}</td>
                    </tr>
                    <tr>
                        <th>Neutral Kills</th>
                        <td>{{ $player['neutral_kills'] }}</td>
                    </tr>
                    <tr>
                        <th>Civilian Kills</th>
                        <td>{{ $player['civilian_kills'] }}</td>
                    </tr>
                    <tr>
                        <th>All Kills</th>
                        <td>{{ \App\Utils::getAllKills($player) }}</td>
                    </tr>
                    <tr>
                        <th>Deaths</th>
                        <td>{{ $player['deaths'] }}</td>
                    </tr>
                    <tr>
                        <th>Clan</th>
                        <td> {{ $player['tag'] }}</td>
                    </tr>
                    @if(!empty($player['packed_past_clans']))
                        <tr>
                            <th>Past Clans</th>
                            {{--TODO Fix style--}}
                            <td>{!! \App\Utils::formatPastClans($player['packed_past_clans']) !!}</td>
                        </tr>
                    @endif
                </table>
            </td>
        </tr>
        </tbody>
    </table>
    <table width="100%">
        <tbody>
        <tr>
            <td width="50%" valign="top">
                <h4>Last 5 Kills</h4>
                <table class="table table-striped">
                    <thead>
                    <tr>
                        <th>Clan</th>
                        <th>Victim</th>
                    </tr>
                    </thead>
                    <tbody>
                    @foreach($player['last_kills'] as $kill)
                        <tr>
                            <td style="cursor: pointer"
                                onclick="showClan('{{$kill->victim_tag}}');">{{$kill->victim_tag}}</td>
                            <td style="cursor: pointer"
                                onclick="showPlayer('{{$kill->victim}}');">{{$kill->victim}}</td>
                        </tr>
                    @endforeach
                    </tbody>
                </table>
            </td>
            <td width="50%" valign="top">
                <h4>Last 5 Deaths</h4>
                <table class="table table-striped">
                    <thead>
                    <tr>
                        <th>Clan</th>
                        <th>Killer</th>
                    </tr>
                    </thead>
                    <tbody>
                    @foreach($player['last_deaths'] as $death)
                        <tr>
                            <td style="cursor: pointer"
                                onclick="showClan('{{$death->attacker_tag}}');">{{$death->attacker_tag}}</td>
                            <td style="cursor: pointer"
                                onclick="showPlayer('{{$death->attacker}}');">{{$death->attacker}}</td>
                        </tr>
                    @endforeach
                    </tbody>
                </table>
            </td>
        </tr>
        </tbody>
    </table>
@endsection
@endisset
