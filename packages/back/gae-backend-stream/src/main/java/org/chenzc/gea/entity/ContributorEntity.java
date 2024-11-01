package org.chenzc.gea.entity;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.experimental.Accessors;

@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
@Accessors(chain = true)
public class ContributorEntity {
    private String login;
    private Integer id;
    private String avatarUrl;
    private String url;
    private String contributions;
}
