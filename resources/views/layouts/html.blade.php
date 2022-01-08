<!doctype html>
<html lang="{{ str_replace('_', '-', app()->getLocale()) }}">
<head>
    <title>@yield('title', "SimpleClanStats")</title>

    <!--  Meta  -->
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <meta name="csrf-token" content="{{ csrf_token() }}">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="theme-color" media="(prefers-color-scheme: light)" content="white">
    <meta name="theme-color" media="(prefers-color-scheme: dark)" content="black">

    <!--  Links  -->
    <link href="{{asset("css/style.css")}}" rel="stylesheet">
    <link rel="icon" href="{{asset("img/favicon.svg")}}">
    <link rel="apple-touch-icon" href="{{asset("img/favicon-mac.png")}}">
    <link rel="manifest" href="{{asset("manifest.json")}}">

    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons&display=swap"
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
