# homelab-dashboard
<img width="1853" height="897" alt="Screenshot from 2025-08-19 03-26-03" src="https://github.com/user-attachments/assets/64a63068-78a1-4c3a-9b53-b5d1fa57bb5d" />


Its a web UI dashboard for your server where you get all these features

- Role based user Login
- See the running docker web-application and which has a label `x-homelab: true` and access them from the UI, these application can be configured based on these labels on there containers
    - `x-homelab-index=<number>` this will be used to put application in order
    - `x-homelab-name`
    - `x-homelab-icon` this will be used to show icon for that application (it should be a web link)
    - `x-homelab-web-url` this will decide on click of the application where will the user go
- Display Server Stats like `Disk` usage, `CPU Usage`, `CPU Temptreture` and `Memory Usage` 

It was built for for `homelabs` built via https://github.com/chiragsoni81245/homelab, but can be used standalone as well.
