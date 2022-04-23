<header>
    <nav>
        <ul>
            <li><a href="{{url('/'.$locale.'/clans')}}">{{ __('messages.header.clans') }}</a></li>
            <li><a href="{{url('/'.$locale)}}">{{ env('APP_NAME') }}</a></li>
            <li><a href="{{url('/'.$locale.'/players')}}">{{ __('messages.header.players') }}</a></li>
        </ul>
    </nav>

    <ul>
        <li><span class="material-icons">brightness_3</span></li>
        <li><span class="material-icons">language</span></li>
    </ul>
</header>
