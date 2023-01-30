$(()=> {
    if (!ensureToken()) {
        return;
    }

    listAlbums(localStorage.getItem('token'));
});


function ensureToken() {
    if (window.debug) {
        localStorage.removeItem('token');
    }

    window.query = {};
    document.location.search.replace('?', '').split('&').forEach((elt => {
        const splitted = elt.split('=');
        if (splitted.length ===2) {
            window.query[decodeURIComponent(splitted[0])] = decodeURIComponent(splitted[1]);
        }
    }))

    if (window.query.token) {
        localStorage.setItem('token', window.query.token);
    }

    if (!localStorage.getItem('token')) {
        location.href = `${window.apiURL}/login`;

        return false;
    }

    return true;
}

function listAlbums(token) {
    $.ajax({
        headers: {
            'Authorization': `Bearer ${token}`,
            'Accept': 'application/json',
        },
        method: 'GET',
        url: `${window.apiURL}/albums`,
        dataType: 'json',
    }).done((data) => {
        data.forEach((elt) => {
            image = elt.images && elt.images.length ? elt.images[0] : {};
            $('#albums').append($(`<div class="album"><img src="${image.url}"><span>${elt.name}</span></div>`))

        });
    });
}