package prived.medved.response;

import lombok.AllArgsConstructor;
import lombok.Data;
import prived.medved.response.enums.EntityType;

import javax.xml.bind.annotation.XmlRootElement;

@Data
@AllArgsConstructor
@XmlRootElement
public class SimpleEntity {
    public EntityType type;
    public String payload;
}
