package processnotifications

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}


func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	}
	return d.priorityLevel
}

func (g groupMessage) importance() int {
	return g.priorityLevel
}


func (s systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	if dm, ok := n.(directMessage); ok {
		return dm.senderUsername, dm.importance()
	} else if gm, ok := n.(groupMessage); ok {
		return gm.groupName, gm.importance()
	} else if sa, ok := n.(systemAlert); ok {
		return sa.alertCode, sa.importance()
	}
	
	return "", 0
}
