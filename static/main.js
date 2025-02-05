// Icons made by Freepik from www.flaticon.com
const BOT_IMG = "http://localhost:8080/static/robot.jpg";
const PERSON_IMG = "http://localhost:8080/static/man.png";
const BOT_NAME = "BOT";
const PERSON_NAME = "Anoop";
new_chat = true;
var msgerForm, msgerInput, msgerChat
document.addEventListener('DOMContentLoaded', function () {
    msgerForm = get(".msger-inputarea");
    msgerInput = get(".msger-input");
    msgerChat = get(".msger-chat");
    msgerForm.addEventListener("submit", event => {
        event.preventDefault();
        const msgText = msgerInput.value;
        if (!msgText) return;
        appendMessage(PERSON_NAME, PERSON_IMG, "right", msgText);
        msgerInput.value = "";
        botResponse(msgText);
    });
});

function appendMessage(name, img, side, text) {
  //   Simple solution for small apps
  var messageId = crypto.randomUUID();
  const msgHTML = `
    <div class="msg ${side}-msg" style="align-items: flex-start">
      <div class="msg-img" style="background-image: url(${img})"></div>

      <div class="msg-bubble">
        <div class="msg-info">
          <div class="msg-info-name">${name}</div>
          <div class="msg-info-time">${formatDate(new Date())}</div>
        </div>

        <div class="msg-text" id="${messageId}">${text}</div>
      </div>
    </div>
  `;

  msgerChat.insertAdjacentHTML("beforeend", msgHTML);
  msgerChat.scrollTop += 500;
  return messageId;
}

function botResponse(msgText) {
  var chatMessage = {
    question: msgText,
    new_chat: new_chat
  };
  var messageId = appendMessage(BOT_NAME, BOT_IMG, "left", "");
  processStreamingResponse('http://localhost:8080/v1/test/stream', messageId, chatMessage)
}

// Utils
function get(selector, root = document) {
  return root.querySelector(selector);
}

function formatDate(date) {
  const h = "0" + date.getHours();
  const m = "0" + date.getMinutes();

  return `${h.slice(-2)}:${m.slice(-2)}`;
}

function random(min, max) {
  return Math.floor(Math.random() * (max - min) + min);
}

async function processStreamingResponse(url, messageId, chatMessage) {
  console.log(messageId);
  const response = await fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(chatMessage)
  });
  new_chat = false;
  if (!response.ok) {
    var textElement = document.getElementById(messageId);
    var messageText = document.createTextNode('Network response was not ok');
    textElement.appendChild(messageText);
    throw new Error('Network response was not ok');
  }
  const reader = response.body.getReader();
  const decoder = new TextDecoder();
  while (true) {
    const { value, done } = await reader.read();
    if (done) {
      break;
    }
    const chunk = decoder.decode(value);
    console.log(chunk);
    // Process the chunk of streaming data
    var textElement = document.getElementById(messageId);
    textElement.innerHTML += chunk;
  }
}
