import React, { useState } from 'react';
import './App.css';

function App() {
    const [key, setKey] = useState('');
    const [keys, getKey] = useState('');
    const [value, setValue] = useState('');
    const [expiration, setExpiration] = useState('');
    const [result, setResult] = useState(null);

    const handleGet = async () => {
        try {
            const response = await fetch(`http://localhost:8080/get_lru_key/${keys}`);
            if (!response.ok) {
                throw new Error('Key not found');
            }
            const data = await response.json();
            setResult(data.value);
        } catch (error) {
            setResult(error.message);
        }
    };

    const handleSet = async () => {
        try {
            const response = await fetch('http://localhost:8080/set', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    key: key,
                    value: value,
                    expiration: Number(expiration)
                })
            });
            if (!response.ok) {
                throw new Error('Error setting key');
            }
            const data = await response.json();
            setResult('Key set successfully');
        } catch (error) {
            setResult(error.message);
        }
    };

    return (
        <div className="App">
            <header className="App-header">
                <h1>Set/Fetch Least Recently Used (LRU) Cache</h1>
                
                <div>
                    <input
                        type="text"
                        placeholder="Key"
                        value={key}
                        onChange={(e) => setKey(e.target.value)}
                    />
                    <input
                        type="text"
                        placeholder="Value"
                        value={value}
                        onChange={(e) => setValue(e.target.value)}
                    />
                    <input
                        type="text"
                        placeholder="Expiration (seconds)"
                        value={expiration}
                        onChange={(e) => setExpiration(e.target.value)}
                    />
                    <button onClick={handleSet}>Set</button>
                </div>
              

                <div>
                    <input
                        type="text"
                        placeholder="Key"
                        value={keys}
                        onChange={(e) => getKey(e.target.value)}
                    />
                    <button onClick={handleGet}>Get</button>
                </div>

                <div>
                    <h2>Result :-  {result}</h2>
                </div>


            </header>
        </div>
    );
}

export default App;
