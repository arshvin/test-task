package prived.medved.response;

import lombok.AllArgsConstructor;
import lombok.Data;

import javax.xml.bind.annotation.XmlRootElement;

@Data
@AllArgsConstructor
@XmlRootElement
public class IFaceIPPair {
    public String name;
    public String ip;
}
