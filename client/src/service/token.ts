export function ensureToken() {
    if ((window as any).debug) {
        localStorage.removeItem('token');
    }

    (window as any).query = {};
    document.location.search.replace('?', '').split('&').forEach((elt => {
        const splitted = elt.split('=');
        if (splitted.length ===2) {
            (window as any).query[decodeURIComponent(splitted[0])] = decodeURIComponent(splitted[1]);
        }
    }))

    if ((window as any).query.token) {
        localStorage.setItem('token', (window as any).query.token);
    }

    if (!localStorage.getItem('token')) {
        document.location.href = `${(window as any).apiURL}/login`;

        return false;
    }

    return true;
}