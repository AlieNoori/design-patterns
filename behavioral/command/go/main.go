package main

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}
	offCommand := &OffCommand{
		device: tv,
	}

	button := &Button{
		command: onCommand,
	}

	button.press()

	button.command = offCommand

	button.press()
}
