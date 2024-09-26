package prived.medved.response;

import lombok.AllArgsConstructor;
import lombok.Data;
import prived.medved.response.enums.EntityType;

@Data
@AllArgsConstructor
public class SimpleEntity {
  public EntityType type;
  public String payload;
}
