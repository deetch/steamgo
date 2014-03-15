/*
This package provides access to Steam as if it was an actual Steam client.

To login, you'll have to create a new Client first. Then connect to the Steam network
and wait for a ConnectedCallback. This means you can now call the Login method in the Auth module
with your login information. This is covered in more detail in the method's documentation. After you've
received the LoggedOnEvent, you should set your persona state to online to receive friend lists etc.

	client := steamgo.NewClient()
	client.Connect()
	for event := range client.Events() {
		switch e := event.(type) {
		case steamgo.ConnectedEvent:
			client.Auth.LogOn(myLoginInfo)
		case steamgo.MachineAuthUpdateEvent:
			ioutil.WriteFile("sentry", e.Hash, 0666)
		case steamgo.LoggedOnEvent:
			client.Social.SetPersonaState(internal.EPersonaState_Online)
		case steamgo.FatalError:
			client.Connect() // please do some real error handling here
			log.Print(e)
		case error:
			log.Print(e)
		}
	}

*/
package steamgo
