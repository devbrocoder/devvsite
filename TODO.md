# Devvsite — Development Roadmap

> A Go-based CLI tool for creating and managing native PHP development environments on Linux.

---

# 0. Project Goals

Devvsite should make local PHP development on Linux simple and repeatable.

Instead of manually doing:

- `mkdir`
- configuring permissions
- creating MySQL databases
- editing `/etc/hosts`
- creating Nginx server blocks
- enabling Nginx sites
- configuring PHP-FPM
- restarting/reloading services
- checking configuration
- troubleshooting common errors

the developer should be able to use commands such as:

```bash
devvsite create myapp
devvsite list
devvsite info myapp
devvsite start myapp
devvsite stop myapp
devvsite restart myapp
devvsite delete myapp
devvsite doctor
