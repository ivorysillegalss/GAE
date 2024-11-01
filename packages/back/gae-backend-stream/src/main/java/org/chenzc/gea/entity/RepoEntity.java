package org.chenzc.gea.entity;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.experimental.Accessors;


/**
 * @author chenz
 * @date 2024/10/26
 */
@Builder
@Data
@AllArgsConstructor
@NoArgsConstructor
@Accessors(chain = true)
public class RepoEntity {
    private int repoId;
    private String ownerName;
    private String ownerId;
    private String name;
    private Integer createdAt;
    private Integer updatedAt;
    private Integer forksCount;
    private Integer networkCount;
    private Integer openIssuesCount;
    private Integer stargazersCount;
    private Integer watchersCount;
    private String[] ContributorsId;
}
