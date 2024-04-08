import  { useState } from 'react';

const App = () => {
  const [key, setKey] = useState('');
  const [value, setValue] = useState('');
  const [result, setResult] = useState('');

  const handleSet = async () => {
    const response = await fetch('http://localhost:5000/cache/set', {
      method: 'POST',
      body: new URLSearchParams({ key, value, expiration: '5s' }),
    });
    const data = await response.text();
    setResult(data);
  };

  const handleGet = async () => {
    const response = await fetch(`http://localhost:5000/cache/get/${key}`);
    const data = await response.text();
    setResult(data);
  };

  return (
    <div>
      <input type="text" placeholder="Key" value={key} onChange={(e) => setKey(e.target.value)} />
      <input type="text" placeholder="Value" value={value} onChange={(e) => setValue(e.target.value)} />
      <button onClick={handleSet}>Set</button>
      <button onClick={handleGet}>Get</button>
      <div>{result}</div>
    </div>
  );
};

export default App;
