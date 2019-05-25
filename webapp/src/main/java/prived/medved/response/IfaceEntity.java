package prived.medved.response;

import lombok.AllArgsConstructor;
import lombok.Data;
import prived.medved.response.enums.EntityType;
import prived.medved.response.enums.IPsCount;

import javax.xml.bind.annotation.XmlRootElement;
import java.util.List;

@Data
@AllArgsConstructor
@XmlRootElement
public class IfaceEntity {
    public EntityType type;
    public IPsCount count;
    public List payload;
}
