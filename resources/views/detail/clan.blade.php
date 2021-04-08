@extends('layouts.modal')
@isset($clan)
    @section('modal-title', $clan["name"])
@section('modal-content')
    {{--    TODO translate messages--}}
    <table class="w-100">
        <tbody>
        <tr>
            <td class="align-top">
                <h3>Members</h3>
                <table class="table table-striped">
                    <tbody>
                    @foreach($clan["members"] as $member)
                        <tr>
                            <td><img src="https://mc-heads.net/avatar/{{ $member }}/16/"
                                     alt="{{ $member }}s Avatar"></td>
                            <td class="modal-opener" data-nick="{{$member}}">{{ $member }}</td>
                        </tr>
                    @endforeach
                    </tbody>
                </table>
            </td>
            <td class="align-top">
                <h3>Info</h3>
                <table class="table">
                    <tbody>
                    <tr>
                        <th>Name</th>
                        <td>{{$clan["name"]}}</td>
                    </tr>
                    <tr>
                        <th>Clan Tag</th>
                        <td>{!! \App\Utils::addColors($clan["color_tag"]) !!}</td>
                    </tr>
                    <tr>
                        <th>Verified</th>
                        @if($clan["verified"])
                            <td>Yes</td>
                        @else
                            <td>No</td>
                        @endif
                    </tr>
                    <tr>
                        <th>Founded</th>
                        <td>{{\App\Utils::formatDate($clan["founded"])}}</td>
                    </tr>
                    <tr>
                        <th>Leaders</th>
                        <td>{{join(", ", $clan["leaders"])}}</td>
                    </tr>
                    <tr>
                        <th>Members</th>
                        <td>{{count($clan["members"])}}</td>
                    </tr>
                    <tr>
                        <td colspan="2"><br></td>
                    </tr>
                    <tr>
                        <th>Balance</th>
                        <td>{{$clan["balance"]}}</td>
                    </tr>
                    <tr>
                        <th>KDR</th>
                        <td>{{number_format($clan["KDR"], 2)}}</td>
                    </tr>
                    <tr>
                        <th>Allies</th>
                        <td>{{ \App\Utils::unpackClans($clan['packed_allies']) }}</td>
                    </tr>
                    <tr>
                        <th>Rivals</th>
                        <td>{{ \App\Utils::unpackClans($clan['packed_rivals'])}}</td>
                    </tr>
                    <tr>
                        <th>Wars</th>
                        <td>{{ \App\Utils::getWarringFromFlags($clan['flags']) }}</td>
                    </tr>
                    </tbody>
                </table>
            </td>
        </tr>
        </tbody>
    </table>
@endsection
@endisset
