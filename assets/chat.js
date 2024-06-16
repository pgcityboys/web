document.addEventListener('DOMContentLoaded', () => {
    const chatMessages = document.getElementById('chat-messages');
    const chatForm = document.getElementById('chat-form');
    const messageInput = document.getElementById('message-input');

    let websocket;

    function getChatUser() {
        const userJSON = localStorage.getItem('user');
        if (userJSON) {
            try {
                return JSON.parse(userJSON);
            } catch (e) {
                console.error('Error parsing user data:', e);
            }
        }
        const uuid = crypto.randomUUID();
        localStorage.setItem('user', JSON.stringify(uuid));
        return uuid;
    }

    // Setup WebSocket connection
    function setupWebSocket() {
        const userId = getChatUser();
        const websocketUrl = `ws://localhost:2137/ws/${userId}`; 

        websocket = new WebSocket(websocketUrl);
        websocket.onopen = () => {
            console.log('Connected to WebSocket server');
        };

        websocket.onmessage = (event) => {
            let [author, ...message] = event.data.split(':');
            message = message.join(':')
            displayMessage(author, message, 'received');
        };

        websocket.onclose = (event) => {
            console.log('WebSocket connection closed');
            displayMessage("", event.data, 'received');
        };

        websocket.onerror = (error) => {
            console.error('WebSocket error:', error);
        };
    }

    // Send message through WebSocket
    function sendMessage(event) {
        event.preventDefault();

        const messageContent = messageInput.value.trim();
        if (messageContent === '') return;
        const message = {
            user: getChatUser(),
            event: 'chat',
            data: {
                room: 'c18abd89-c55d-4642-a1fc-1af701e8411b',
                content: messageContent
            }
        };

        websocket.send(JSON.stringify(message));
        displayMessage('You', messageContent, 'sent');
        messageInput.value = '';
    }

    // Display message in chat
    function displayMessage(author, message, type) {
        const messageElement = document.createElement('div');
        messageElement.classList.add('message', type);
        if (author == getChatUser() || author == '') {
            return;
        }
        messageElement.textContent = `${author}: ${message}`;
        chatMessages.appendChild(messageElement);
        chatMessages.scrollTop = chatMessages.scrollHeight;
    }

    // Initialize WebSocket connection and form submission handler
    setupWebSocket();
    chatForm.addEventListener('submit', sendMessage);
});
