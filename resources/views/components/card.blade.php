<div class="card">
    <div class="card-body">
        <h4>{{ $name }}</h4>
        <h5><span class="material-icons">people</span>{{ $members }}</h5>
    </div>
    <div class="card-footer">
        <span>{{$score}}</span>
        <span>
            @for($i = 0; $i < $stars; $i++)
                <span class="material-icons">star</span>
            @endfor
        </span>
    </div>
</div>
