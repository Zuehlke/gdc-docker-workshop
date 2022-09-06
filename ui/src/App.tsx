import React, {useState} from 'react';
import './App.css';

function App() {
    const [releaseTitle, setReleaseTitle] = useState('First release!')
    const [releaseVersion, setReleaseVersion] = useState('v0.0.1')
    const [releaseDate, setReleaseDate] = useState('2022-08-10T00:00:00Z')

    const submitRelease = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        const requestBody = {
            releaseTitle: releaseTitle,
            releaseVersion: releaseVersion,
            releaseDate: releaseDate
        }
        console.log(`Submitting release: ${JSON.stringify(requestBody, null, 2)}`)

        // TODO docker-compose setup
        const response = await fetch('http://backend/releases', {
            method: 'POST',
            headers: {
                'AUTH_TOKEN': 'TODO'
            },
            body: JSON.stringify(requestBody)
        })
        const responseBody = await response.json()
        console.log(`Received response ${JSON.stringify(responseBody, null, 2)}`)
        alert(`Received response ${JSON.stringify(responseBody, null, 2)}`)
    }

    return (
        <div className="App">
            <header className="App-header">
                <form onSubmit={(e) => submitRelease(e)}>
                    <label>Release title:</label><br/>
                    <input type="text" id="releaseTitle" name="releaseTitle" value={releaseTitle} onChange={(e) => setReleaseTitle(e.target.value)}/><br/>
                    <label>Release version: - (example: v0.0.0)</label><br/>
                    <input type="text" id="releaseVersion" name="releaseVersion" value={releaseVersion} onChange={(e) => setReleaseVersion(e.target.value)}/><br/>
                    <label>Release date: - (ISO format)</label><br/>
                    <input type="text" id="releaseDate" name="releaseDate" value={releaseDate} onChange={(e) => setReleaseDate(e.target.value)}/><br/>
                    <input type="submit" value="Schedule release!"/>
                </form>
                <h6>(UX/UI you never seen before!)</h6>
            </header>
        </div>
    );
}

export default App;
