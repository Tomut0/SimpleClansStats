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
    <link href="{{asset("css/app.css")}}" rel="stylesheet">
    <link rel="icon" href="{{asset("img/favicon.svg")}}">
    <link rel="apple-touch-icon" href="{{asset("img/favicon-mac.png")}}">
    <link rel="manifest" href="{{asset("manifest.json")}}">
    <meta name="theme-color" media="(prefers-color-scheme: light)" content="white">
    <meta name="theme-color" media="(prefers-color-scheme: dark)" content="black">

    @stack('scripts')
</head>
<body>
@include('layouts.header')
@yield('content')
@include('layouts.footer')
<script src="{{asset("js/app.js")}}"></script>
</body>
</html>
