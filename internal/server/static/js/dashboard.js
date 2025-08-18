// Toggle dropdown on profile button click
const profileBtn = document.getElementById("profileBtn");
const dropdownMenu = document.getElementById("dropdownMenu");

profileBtn.addEventListener("click", () => {
    dropdownMenu.classList.toggle("hidden");
});

async function updateSystemStatus() {
    let [ok, data] = await _fetch("/system/stats");

    if (!ok) {
        showToast(data.error);
        return;
    }

    const cpuBar = document.getElementById("cpuBar");
    const cpuUsage = document.getElementById("cpuUsage");
    const cpuTempreture = document.getElementById("cpuTempreture");

    cpuBar.style.width = `${data.cpu.usage.toFixed(2)}%`;
    cpuUsage.textContent = `${data.cpu.usage.toFixed(2)}%`;
    if (data.cpu.tempreture > 0) {
        cpuTempreture.classList.remove("hidden");
    } else {
        cpuTempreture.classList.add("hidden");
    }
    cpuTempreture.textContent = `${Math.floor(data.cpu.tempreture)}°C`;

    const memoryBar = document.getElementById("memoryBar");
    const memoryUsage = document.getElementById("memoryUsage");
    const memoryText = document.getElementById("memoryText");
    memoryBar.style.width = `${data.memory.usage.toFixed(2)}%`;
    memoryUsage.textContent = `${data.memory.usage.toFixed(2)}%`;
    memoryText.textContent = `${(data.memory.used / Math.pow(10, 9)).toFixed(2)} / ${Math.floor(data.memory.total / Math.pow(10, 9))}`;

    // For disks, you can extend to dynamic rendering
    const getDiskTemplate = ({ name, usage, total }) =>
        `
<div>
    <span class="text-md text-white-700">${capitalize(name)}</span>
    <div class="w-full bg-gray-200 rounded-full h-3">
        <div class="bg-purple-500 h-3 rounded-full" style="width: ${usage}%"></div>
    </div>
    <div class="flex flex-row justify-between items-baseline">
        <p class="text-sm text-white-300 m-0 p-0">${usage.toFixed(2)}%</p>
        <p class="text-xs text-white-300 m-0 p-0">${((total * usage) / 100).toFixed(2)}GB / ${Math.floor(total)}GB</p>
    </div>
</div>`.trim();

    const disksContainer = document.getElementById("disks-container");
    disksContainer.innerHTML = "";
    for (let disk of data.disks || []) {
        let diskElement = getTemplateToElement(getDiskTemplate(disk));
        disksContainer.appendChild(diskElement);
    }
}

async function loadApplications() {
    let [ok, data] = await _fetch("/applications");
    if (!ok) {
        showToast(data.error);
        return;
    }

    const getApplicationTemplate = ({ name, web_url, icon }) =>
        `
<a href="${web_url}" target="_blank" class="lg:md:w-40 lg:md:p-4 my-2 lg:md:my-0 w-30 p-2 bg-black/30 text-white rounded-lg shadow hover:shadow-lg transition flex flex-col items-center">
    <img src="${icon}" alt="App" class="w-16 h-16 mb-2">
    <span class="font-medium">${capitalize(name)}</span>
</a>`.trim();

    const applicationsContainer = document.getElementById(
        "applications-container"
    );
    applicationsContainer.innerHTML = "";
    for (let application of data || []) {
        let applicationElement = getTemplateToElement(
            getApplicationTemplate(application)
        );
        applicationsContainer.appendChild(applicationElement);
    }
}

function main() {
    loadApplications();
    updateSystemStatus();
    setInterval(
        updateSystemStatus,
        parseInt(document.getElementById("update_frequency").value) || 5000
    );
}
main();
