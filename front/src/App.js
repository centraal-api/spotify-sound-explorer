import React, { useState, useEffect } from 'react';

function App() {
  const [message, setMessage] = useState('');

  const fetchMessage = () => {
    fetch('https://spotify-sound-explorer-cc4f60bc745e.herokuapp.com/get') // Replace with your backend API endpoint
      .then(response => response.json())
      .then(data => setMessage(data.message))
      .catch(error => console.error('Error fetching data:', error));
  };

  return (
    <div>
      <h1>Message from Backend:</h1>
      <p>{message}</p>
      <button onClick={fetchMessage}>Fetch Message</button>
    </div>
  );
}

export default App;
