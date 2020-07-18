package prived.medved.response;

import lombok.AllArgsConstructor;
import lombok.Data;
import prived.medved.response.enums.EntityType;
import prived.medved.response.enums.IPsCount;
import java.util.List;

@Data
@AllArgsConstructor
public class IfaceEntity {
    public EntityType type;
    public IPsCount count;
    public List payload;
}
