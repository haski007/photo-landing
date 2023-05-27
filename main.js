// This would ideally be fetched from a server
const portfolioImages = [
    "resources/style.jpg",
    "resources/batman.jpg",
    "resources/princess.jpg",
    "resources/Joint.jpg",
    "resources/sad.jpg",
    "resources/serious.jpg",
// ...
];
AOS.init();

// Populate portfolio images
const portfolioDiv = document.getElementById("portfolio");
portfolioImages.forEach((img) => {
    const imgDiv = document.createElement("div");
    imgDiv.classList.add("col-md-4");
    imgDiv.innerHTML = `<img src="${img}" class="img-fluid">`;
    portfolioDiv.appendChild(imgDiv);
});


// Change language
function changeLang(lang) {
    document.querySelector("#title h1").innerText = langData[lang].title;

    const contactHeading = document.getElementById("contact-heading");
    const nameLabel = document.getElementById("name-label");
    const emailLabel = document.getElementById("email-label");
    const messageLabel = document.getElementById("message-label");
    const submitButton = document.getElementById("submit-button");

    const nameInput = document.getElementById("name");
    const emailInput = document.getElementById("email");
    const messageInput = document.getElementById("message");

    contactHeading.innerText = langData[lang].contact.heading;
    nameLabel.innerText = langData[lang].contact.name;
    emailLabel.innerText = langData[lang].contact.email;
    messageLabel.innerText = langData[lang].contact.message;
    submitButton.innerText = langData[lang].contact.button;

    nameInput.placeholder = langData[lang].contact.placeholders.name;
    emailInput.placeholder = langData[lang].contact.placeholders.email;
    messageInput.placeholder = langData[lang].contact.placeholders.message;
}


document.getElementById('contact-form').addEventListener('submit', function(e) {
    e.preventDefault();

    // Get the input values
    let name = document.getElementById('name').value;
    let email = document.getElementById('email').value;
    let message = document.getElementById('message').value;

    // Display a simple alert message
    alert(`Thanks for your message, ${name}! We'll get back to you at ${email} as soon as possible.`);
});

document.addEventListener("DOMContentLoaded", function () {
});

// Language data
const langData = {
    pl: {
        title: "Portfolio Fotograficzne",
        contact: {
            heading: "Kontakt",
            name: "Imię",
            email: "Email",
            message: "Wiadomość",
            button: "Wyślij",
            placeholders: {
                name: "Twoje imię",
                email: "twój.email@example.com",
                message: "Twoja wiadomość"
            }
        }
    },
    en: {
        title: "Photography Portfolio",
        contact: {
            heading: "Contact",
            name: "Name",
            email: "Email",
            message: "Message",
            button: "Send",
            placeholders: {
                name: "Your name",
                email: "your.email@example.com",
                message: "Your message"
            }
        }
    },
    ua: {
        title: "Портфоліо фотографії",
        contact: {
            heading: "Контакт",
            name: "Ім'я",
            email: "Електронна пошта",
            message: "Повідомлення",
            button: "Надіслати",
            placeholders: {
                name: "Ваше ім'я",
                email: "ваш.email@example.com",
                message: "Ваше повідомлення"
            }
        }
    }
};
