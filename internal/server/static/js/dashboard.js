// Toggle dropdown on profile button click
const profileBtn = document.getElementById("profileBtn");
const dropdownMenu = document.getElementById("dropdownMenu");

profileBtn.addEventListener("click", () => {
    dropdownMenu.classList.toggle("hidden");
});

function updateSystemStatus(cpu, memory, disks) {
    const cpuBar = document.getElementById("cpuBar");
    const cpuText = document.getElementById("cpuText");
    cpuBar.style.width = `${cpu}%`;
    cpuText.textContent = `${cpu}%`;

    const memoryBar = document.getElementById("memoryBar");
    const memoryText = document.getElementById("memoryText");
    memoryBar.style.width = `${memory}%`;
    memoryText.textContent = `${memory}%`;

    // For disks, you can extend to dynamic rendering
}

// Example dummy update
setTimeout(() => {
    updateSystemStatus(55, 72, []);
}, 2000);
