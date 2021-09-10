@extends('layouts.html')
@section('title', env("APP_NAME"))
@section('content')
    <main class="mt-3">
        <div class="container-fluid">
            <div class="row text-center">
                <div class="col">
                    <div class="card">
                        <div class="card-header bg-primary">
                            <div class="container">
                                <div class="row">
                                    <div class="col-sm">
                                        <!-- Todo: add translation for card title -->
                                        <h4 class="p-3 mb-0 text-white">Clan List</h4>
                                    </div>
                                    <div class="col-sm">
                                        <form class="d-flex p-3 w-75 float-end">
                                            <input id="filter" class="form-control" type="search"
                                                   placeholder="Filter"
                                                   aria-label="Filter">
                                        </form>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="card-body">
                            <table class="table w-100">
                                <thead>
                                <tr>
                                    <th class="col" scope="col"># <i class="fas fa-sort"></i></th>
                                    <th class="col" scope="col">{{ __('messages.clans.clanTag') }} <i
                                            class="fas fa-sort"></i></th>
                                    <th class="col-2" scope="col">{{ __('messages.main.clanName') }} <i
                                            class="fas fa-sort"></i></th>
                                    <th class="col-2" scope="col">{{ __('messages.main.clanKDR') }} <i
                                            class="fas fa-sort"></i></th>
                                    <th class="col-1" scope="col">{{ __('messages.clans.clanLeaders') }} <i
                                            class="fas fa-sort"></i>
                                    </th>
                                    <th class="col-2" scope="col">{{ __('messages.main.clanMembers') }} <i
                                            class="fas fa-sort"></i>
                                    </th>
                                    <th class="col-1" scope="col">{{ __('messages.clans.clanFounded') }} <i
                                            class="fas fa-sort"></i>
                                    </th>
                                    <th class="col-2" scope="col">{{ __('messages.clans.clanLastUsed') }} <i
                                            class="fas fa-sort"></i>
                                    </th>
                                </tr>
                                </thead>
                                <tbody>
                                @foreach($clans->getTopClans(999999, "KDR") as $index => $clan)
                                    @if(empty($clan['name']) || !array_key_exists('color_tag', $clan))
                                        @continue
                                    @endif
                                    <tr class="data_row">
                                        <th scope="row">{{ $clan['rank'] }}</th>
                                        <td>{!! $clan['color_tag'] !!}</td>
                                        <td class="modal-opener clan_name" data-tag='{{$clan["tag"]}}'>{{ $clan['name'] }}</td>
                                        <td>{{ number_format($clan['KDR'], 2) }}</td>
                                        <td>{{ sizeof($clan['leaders']) }}</td>
                                        <td>{{ sizeof($clan['members']) }}</td>
                                        <td> {{ Utils::formatDate($clan['founded']) }}</td>
                                        <td>{{ Utils::formatDateTime($clan['last_used']) }}</td>
                                    </tr>
                                @endforeach
                                </tbody>
                            </table>
                        </div>
                        <div class="card-footer">
                            <nav aria-label="Page navigation">
                                <ul class="pagination justify-content-end">
                                    <li class="page-item"><a class="page-link" href="#">Previous</a></li>
                                    <li class="page-item"><a class="page-link" href="#">1</a></li>
                                    <li class="page-item"><a class="page-link" href="#">2</a></li>
                                    <li class="page-item"><a class="page-link" href="#">3</a></li>
                                    <li class="page-item"><a class="page-link" href="#">Next</a></li>
                                </ul>
                            </nav>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </main>
@endsection
