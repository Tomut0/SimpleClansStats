<!doctype html>
<html lang="{{ str_replace('_', '-', app()->getLocale()) }}">
<head>
    <title>@yield('title', "SimpleClanStats")</title>

    <!--  Meta  -->
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <meta name="csrf-token" content="{{ csrf_token() }}">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <!--  Links  -->
    <link href="{{asset("css/style.css")}}" rel="stylesheet">
    <link rel="shortcut icon" href="{{asset("img/favicon.png")}}" type="image/png">
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons"
          rel="stylesheet">

    @stack('scripts')
</head>
<body>
@include('layouts.header')
@yield('content')
@include('layouts.footer')
<script src="{{asset("js/app.js")}}"></script>
</body>
</html>
