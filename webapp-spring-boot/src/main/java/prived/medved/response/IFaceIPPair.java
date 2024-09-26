package prived.medved.response;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class IFaceIPPair {
  public String name;
  public String ip;
}
