document.getElementById("loginForm").addEventListener("submit", async (e) => {
    e.preventDefault();

    const username = document.getElementById("username").value.trim();
    const password = document.getElementById("password").value.trim();
    const messageEl = document.getElementById("loginMessage");

    messageEl.classList.add("hidden");

    let [ok, data] = await _fetch(`${BASE_API_URL}/auth/login`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
    });

    if (ok) {
        // You can handle token or redirect here
        window.location.href = "/";
    } else {
        showToast(data.error, "error");
    }
});
