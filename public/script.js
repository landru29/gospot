$(()=> {
    ensureToken();


});


function ensureToken() {
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
        location.href = window.loginURL;
    }
}