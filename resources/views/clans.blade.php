@extends('layouts.html')
@section('title', env("APP_NAME"))
@section('content')
    <main class="mt-3">
        <div class="container-fluid">
            <div class="row text-center">
                <div class="col">
                    <h4 class="bg-primary text-white p-3 mb-0">Clan List</h4>
                    <table class="table table-primary table-striped w-100">
                        <thead>
                        <tr>
                            {{--                            <th scope="col">#</th>--}}
                            <th scope="col">{{ __('messages.clans.clanTag') }}</th>
                            <th scope="col">{{ __('messages.main.clanName') }}</th>
                            <th scope="col">{{ __('messages.main.clanKDR') }}</th>
                            <th scope="col">{{ __('messages.clans.clanLeaders') }}</th>
                            <th scope="col">{{ __('messages.main.clanMembers') }}</th>
                            <th scope="col">{{ __('messages.clans.clanFounded') }}</th>
                            <th scope="col">{{ __('messages.clans.clanLastUsed') }}</th>
                        </tr>
                        </thead>
                        <tbody>
                        @foreach($clans->getTopClans(100, "KDR") as $index => $clan)
                            @if(empty($clan['name']) || !array_key_exists('color_tag', $clan))
                                @continue
                            @endif
                            <tr>
                                {{--                                <th scope="row">{{ $index + 1 }}</th>--}}
                                <td>{!! $clan['color_tag'] !!}</td>
                                <td>{{ $clan['name'] }}</td>
                                <td>{{ number_format($clan['KDR'], 2) }}</td>
                                <td>{{ sizeof($clan['leaders']) }}</td>
                                <td>{{ sizeof($clan['members']) }}</td>
                                <td> {{ date("d/m/Y",$clan['founded'] / 1000) }}</td>
                                <td>{{ date("d/m/Y - H:i", $clan['last_used'] / 1000) }}</td>
                            </tr>
                        @endforeach
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </main>
@endsection
