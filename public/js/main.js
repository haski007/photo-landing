let hostURL;

// check if it's localhost or a domain
if (window.location.hostname === "localhost" && window.location.hostname === "127.0.0.1") {
    hostURL = 'http://localhost:8080';
} else {
    hostURL = "http://161.35.150.170:80";
}

// This would ideally be fetched from a cmd
const portfolioImages = [
    "/static/resources/style.jpg",
    "/static/resources/batman.jpg",
    "/static/resources/princess.jpg",
    "/static/resources/Joint.jpg",
    "/static/resources/sad.jpg",
    "/static/resources/serious.jpg",
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
    currentLang = lang
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

let currentLang = 'pl'; // default language

function getCurrentLanguage() {
    // You can modify this function to get the language from another source.
    return currentLang;
}


document.getElementById('contact-form').addEventListener('submit', function (e) {
    e.preventDefault();
    console.log("hello there")

    // Get the input values
    let name = document.getElementById('name').value;
    let email = document.getElementById('email').value;
    let message = document.getElementById('message').value;

    fetch(hostURL + '/contact', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            name: name,
            email: email,
            message: message
        }),
    })
        .then(response => response.json())
        .then(data => {
            console.log('Success:', data);
            // Get the current language and corresponding notification message
            let lang = getCurrentLanguage();
            let notificationTemplate = langData[lang].notification.success;

            // Replace placeholders in the template with actual values
            let notificationText = notificationTemplate
                .replace('{name}', name)
                .replace('{email}', email);

            // Create a notification
            let notification = document.getElementById('notification');
            notification.innerText = notificationText;
            notification.classList.add('show');
        })
        .catch((error) => {
            console.error('Error:', error);
            let lang = getCurrentLanguage();
            let notificationTemplate = langData[lang].notification.fail;

            // Replace placeholders in the template with actual values
            let notificationText = notificationTemplate
                .replace('{name}', name)
                .replace('{email}', email);

            // Create a notification
            let notification = document.getElementById('notification');
            notification.innerText = notificationText;
            notification.classList.add('show');
        });


    // Get the current language and corresponding notification message
    let lang = getCurrentLanguage();
    let notificationText = langData[lang].notification.success;

    // Create a notification
    let notification = document.getElementById('notification');
    notification.innerText = notificationText;
    notification.classList.add('show');

    // Optional: hide the notification after a few seconds
    setTimeout(function() {
        notification.classList.remove('show');
    }, 5000);
});

document.addEventListener("DOMContentLoaded", function () {
});

window.addEventListener('scroll', function () {
    const contactUsTop = document.getElementById('contact-section').offsetTop;
    const scrollButton = document.getElementById('contact-button');

    // Adjust this value based on the height of your navbar or any other offset
    const offset = 100;

    if (window.scrollY >= (contactUsTop - offset)) {
        // If we've reached the Contact Us block, start fading out the button
        scrollButton.style.opacity = 1 - (window.scrollY - (contactUsTop - offset)) / offset;
    } else {
        // If we're above the Contact Us block, make sure the button is fully visible
        scrollButton.style.opacity = 1;
    }
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
        },
        notification: {
            success: "Dziękujemy za wiadomość, {name}! Skontaktujemy się z Tobą na {email} tak szybko jak to możliwe.",
            fail: "Coś poszło nie tak. Spróbuj ponownie później."
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
        },
        notification: {
            success: "Thanks for your message, {name}! We'll get back to you at {email} as soon as possible.",
            fail: "Something went wrong. Please try again later."
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
        },
        notification: {
            success: "Дякуємо за повідомлення, {name}! Ми зв'яжемося з вами за адресою {email} якомога швидше.",
            fail: "Щось пішло не так. Будь ласка, спробуйте пізніше."
        }
    }
};
