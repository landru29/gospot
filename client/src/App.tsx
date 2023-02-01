import React from 'react';
import logo from './logo.svg';
import './App.css';
import { ensureToken } from './service/token';
import { SpotifyService } from './service/spotify';
import { Album } from './model/album';

interface Props  {
}

interface State {
  albums: Album[]
}
class App extends React.PureComponent<Props, State> {
  componentDidMount() {
    ensureToken();
    SpotifyService.album(localStorage.getItem('token')||'').subscribe({
      next: (albums: Album[]) => {
        this.setState({albums});
      }
    })
  }


  render(): JSX.Element | null {  
    const albums = (this.state || {}).albums || []; 
    return <div className="App">
    <header className="App-header">
      <h1>Album list</h1>
    </header>
    <div>
      {albums.map((album: Album)=><div>
        <div className="album">
          <img src={album.images[0].url}></img>
          <span>{album.name}</span></div>
      </div>)}
    </div>
  </div>
  }
}

export default App;
