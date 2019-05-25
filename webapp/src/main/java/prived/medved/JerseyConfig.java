package prived.medved;

import org.glassfish.jersey.server.ResourceConfig;
import org.springframework.stereotype.Component;
import prived.medved.controllers.EntryEndPoint;

@Component
public class JerseyConfig extends ResourceConfig {

    public JerseyConfig(){
         register(EntryEndPoint.class);
    }
}
