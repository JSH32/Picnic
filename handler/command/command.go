package command

type Handler struct {
	Commands map[string]Command
}

func (h *Handler) GetCommand(invoke string) (Command, bool) {
	cmd, ok := h.Commands[invoke]
	return cmd, ok
}

func (h *Handler) GetCommandListLen() int {
	return len(h.Commands)
}
