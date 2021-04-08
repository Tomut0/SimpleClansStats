<div class="modal fade" id="detailModal" tabindex="-1" aria-labelledby="detailModalLabel" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="detailModalLabel">@yield('modal-title', __("messages.detail.error"))</h5>
                <button id="modalClose" type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>

            <div class="modal-body">
                @hasSection('modal-content')
                    @yield('modal-content')
                @else
                    <div class="alert alert-danger" role="alert">
                        {{__("messages.detail.dataNotFound")}}
                    </div>
                @endif
            </div>
        </div>
    </div>
</div>
