import React, { useState, useEffect } from 'react';

function App() {
  const [message, setMessage] = useState('');

  useEffect(() => {
    fetch('https://spotify-sound-explorer-cc4f60bc745e.herokuapp.com/get') // Replace with your backend API endpoint
      .then(response => response.json())
      .then(data => setMessage(data.message))
      .catch(error => console.error('Error fetching data:', error));
  }, []); // Empty dependency array ensures this effect runs once after initial render

  return (
    <div>
      <h1>Message from Backend:</h1>
      <p>{message}</p>
    </div>
  );
}

export default App;
