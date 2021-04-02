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
    <link href="{{asset("assets/css/styles.css")}}" rel="stylesheet">
    <link rel="shortcut icon" href="{{asset("assets/img/favicon.png")}}" type="image/png">

    <!--  Bootstrap5 CSS  -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-BmbxuPwQa2lc/FVzBcNJ7UAyJxM6wuqIj61tLrc4wSX0szH/Ev+nYRRuWlolflfl" crossorigin="anonymous">

    <script src="@yield('js')"></script>
</head>
<body>
@include('layouts.header')
@yield('content')
@include('layouts.footer')
</body>
</html>
