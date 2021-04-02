@extends('layouts.html')
@section('title', env("APP_NAME"))
@section('content')
    <main class="mt-3">
        <div class="container-fluid">
            <div class="row text-center">
                <div class="col">
                    <div class="card">
                        <div class="card-header bg-primary">
                            <h4 class="p-3 mb-0 text-white">Clan List</h4>
                            <form class="d-flex">
                                <input id="filter" class="form-control mb-3 w-25" type="search" placeholder="Filter"
                                       aria-label="Filter">
                            </form>
                        </div>
                        <div class="card-body">
                            <table class="table w-100">
                                <thead>
                                <tr>
                                    <th scope="col"># <i class="fas fa-sort"></i></th>
                                    <th scope="col">{{ __('messages.clans.clanTag') }} <i class="fas fa-sort"></i></th>
                                    <th scope="col">{{ __('messages.main.clanName') }} <i class="fas fa-sort"></i></th>
                                    <th scope="col">{{ __('messages.main.clanKDR') }} <i class="fas fa-sort"></i></th>
                                    <th scope="col">{{ __('messages.clans.clanLeaders') }} <i class="fas fa-sort"></i>
                                    </th>
                                    <th scope="col">{{ __('messages.main.clanMembers') }} <i class="fas fa-sort"></i>
                                    </th>
                                    <th scope="col">{{ __('messages.clans.clanFounded') }} <i class="fas fa-sort"></i>
                                    </th>
                                    <th scope="col">{{ __('messages.clans.clanLastUsed') }} <i class="fas fa-sort"></i>
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
            </div>
        </div>
    </main>
@endsection
@push('scripts')
    <script>
        function filter(rows, text) {
            outter:
                for (const row of rows) {
                    for (const column of row.children) {
                        if (column.innerText.toLowerCase().includes(text.toLowerCase())) {
                            row.hidden = false;
                            continue outter;
                        }
                    }
                    row.hidden = true;
                }
        }

        let filterInput = document.getElementById("filter");
        filterInput.addEventListener("input", function () {
            let text = filterInput.value;
            const rows = document.getElementsByClassName("data_row");
            filter(rows, text);
        });
    </script>
@endpush
