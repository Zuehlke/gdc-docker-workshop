import React from 'react';
import './App.css';

function App() {

    return (
        <div className="App">
            <header className="App-header">
                <form>
                    <label>Release title:</label><br/>
                    <input type="text" id="releaseTitle" name="releaseTitle" value="First release!"/><br/>
                    <label>Release version: - (example: v0.0.0)</label><br/>
                    <input type="text" id="releaseVersion" name="releaseVersion" value="v0.0.1"/><br/>
                    <label>Release date: - (ISO format)</label><br/>
                    <input type="text" id="releaseDate" name="releaseDate" value="2022-08-10T00:00:00Z"/><br/>
                    <input type="submit" value="Schedule release!"/>
                </form>
                <h6>(UX/UI you never seen before!)</h6>
            </header>
        </div>
    );
}

export default App;
