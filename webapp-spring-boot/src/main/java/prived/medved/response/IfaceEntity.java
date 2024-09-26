package prived.medved.response;

import java.util.List;
import lombok.AllArgsConstructor;
import lombok.Data;
import prived.medved.response.enums.EntityType;
import prived.medved.response.enums.IPsCount;

@Data
@AllArgsConstructor
public class IfaceEntity {
  public EntityType type;
  public IPsCount count;
  public List payload;
}
