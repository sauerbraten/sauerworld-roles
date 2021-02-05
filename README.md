# Sauerworld Discord Roles Bot

Listens for reactions on specific messages and toggles roles on the reacting users.

- 1 message, 1 role: each message maps to exactly one role
- adding any reaction to a message will assign the role (if the user doesn't already have it)
- removing any reaction from that message will remove the role (if the user has it)

## Configuration

Configuration is done via environment variables:

- `TOKEN`: Discord bot token
- `ROLES_BY_MESSAGE_ID`: list of pairs of a message ID and a role name each, for example: `ROLES_BY_MESSAGE_ID=492453791002525697=duel,492453832035532800=mix`

Role names are resolved to role IDs via the Discord API and then cached.

Changing or adding roles on the server or mapping roles to new messages requires a restart of the bot with updated configuration.
