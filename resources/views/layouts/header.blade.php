<header>
    <div class="container-fluid">
        <nav class="navbar navbar-expand-lg navbar-light bg-light">
            <div class="collapse navbar-collapse">
                <div class="container-fluid">
                    <ul class="navbar-nav navbar-nav me-auto mb-2 mb-lg-0">
                        <!--suppress HtmlUnknownTag -->
                        <a class="navbar-brand" href="#">{{ env('APP_NAME') }}</a>
                        <li class="navbar-item"><a href="{{url('/'.$locale)}}" class="nav-link active">{{ __('messages.header.home') }}</a></li>
                        <li class="navbar-item"><a href="{{url('/'.$locale.'/clans')}}" class="nav-link">{{ __('messages.header.clans') }}</a></li>
                        <li class="navbar-item"><a href="{{url('/'.$locale.'/players')}}" class="nav-link">{{ __('messages.header.players') }}</a></li>
                    </ul>
                </div>
            </div>
        </nav>
    </div>
</header>
