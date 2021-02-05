package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/sauerbraten/sauerworld-roles/config"
)

func toggleRole(s *discordgo.Session, mr *discordgo.MessageReaction, on bool) {
	// check what role should be toggled
	role, ok := config.RolesByMessageID[mr.MessageID]
	if !ok {
		// not a message we're watching
		return
	}

	roleID, err := getRoleID(s, mr.GuildID, role)
	if err != nil {
		log.Printf("resolving role name '%s' to ID: %v\n", role, err)
		return
	}

	// check if user already has that role
	member, err := s.GuildMember(mr.GuildID, mr.UserID)
	if err != nil {
		log.Printf("looking up reaction user %s in guild %s: %v\n", mr.UserID, mr.GuildID, err)
		return
	}
	found := false
	for _, userRole := range member.Roles {
		if userRole == roleID {
			found = true
			break
		}
	}
	if found == on {
		return
	}

	memberName := fmt.Sprintf("%s#%s [%s]", member.User.Username, member.User.Discriminator, member.User.ID)
	if member.Nick != "" {
		memberName = fmt.Sprintf("%s (%s#%s) [%s]", member.Nick, member.User.Username, member.User.Discriminator, member.User.ID)
	}

	if on {
		log.Printf("assigning role %s [%s] to %s\n", role, roleID, memberName)
		err = s.GuildMemberRoleAdd(mr.GuildID, mr.UserID, roleID)
	} else {
		log.Printf("removing role %s [%s] from %s\n", role, roleID, memberName)
		err = s.GuildMemberRoleRemove(mr.GuildID, mr.UserID, roleID)
	}
	if err != nil {
		log.Printf("error toggling role %s [%s] for user %s (guild %s): %v\n", role, roleID, memberName, mr.GuildID, err)
	}
}

var roleIDsByName = map[string]string{}

func getRoleID(s *discordgo.Session, guildID, name string) (string, error) {
	name = strings.ToLower(name)

	// check cache first
	if id, ok := roleIDsByName[name]; ok {
		return id, nil
	}

	roles, err := s.GuildRoles(guildID)
	if err != nil {
		return "", fmt.Errorf("getting all roles from Discord: %w", err)
	}

	for _, role := range roles {
		if strings.ToLower(role.Name) == name {
			roleIDsByName[name] = role.ID // cache for next time
			return role.ID, nil
		}
	}

	return "", fmt.Errorf("a role with name '%s' does not exist", name)
}
