let userName = "ClientID";
let password = "secretKey";

function authUser(user, password) {
    const token = user + ":" + password;
    const hash = btoa(token);
    return "Basic " + hash;
}

function callWebAPI() {
    let request = new XMLHttpRequest();
    
    // Gunakan JSONPlaceholder untuk testing
    request.open("GET", "https://jsonplaceholder.typicode.com/posts/1", true);
    request.setRequestHeader("Authorization", authUser(userName, password));
    
    request.onreadystatechange = function() {
        if (request.readyState === 4) {
            alert("Status: " + request.status);
            console.log(request.responseText);
            document.getElementById('response').innerHTML = `
                <h3>Status: ${request.status}</h3>
                <pre>${request.responseText}</pre>
            `;
        }
    };
    
    request.send();
}