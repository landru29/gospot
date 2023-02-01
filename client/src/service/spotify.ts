import { map } from 'rxjs/operators';
import { AjaxResponse, ajax } from 'rxjs/ajax';
import { Observable } from 'rxjs';
import { Album } from '../model/album';


export class SpotifyService {
    static album(token: string): Observable<Album[]> {
        return ajax({
            url: `${(window as any).apiURL}/albums`,
            headers: {
                'Authorization': `Bearer ${token}`,
                'Accept': 'application/json',
            },
            method: 'GET',
        }).pipe(
            map((response: AjaxResponse<any>) => {
                return response.response.map((elt: any) => new Album(elt))
            })
        );
    }
}