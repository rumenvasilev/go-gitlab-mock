package mock

// Code generated; DO NOT EDIT.

var GetApiV4AdminBatchedBackgroundMigrations EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/admin/batched_background_migrations",
	Method:  "GET",
}

var GetApiV4AdminBatchedBackgroundMigrationsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/admin/batched_background_migrations/{id}",
	Method:  "GET",
}

var PutApiV4AdminBatchedBackgroundMigrationsPauseById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/admin/batched_background_migrations/{id}/pause",
	Method:  "PUT",
}

var PutApiV4AdminBatchedBackgroundMigrationsResumeById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/admin/batched_background_migrations/{id}/resume",
	Method:  "PUT",
}

var GetApiV4AdminCiVariables EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/admin/ci/variables",
	Method:  "GET",
}

var PostApiV4AdminCiVariables EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/admin/ci/variables",
	Method:  "POST",
}

var DeleteApiV4AdminCiVariablesByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/admin/ci/variables/{key}",
	Method:  "DELETE",
}

var GetApiV4AdminCiVariablesByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/admin/ci/variables/{key}",
	Method:  "GET",
}

var PutApiV4AdminCiVariablesByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/admin/ci/variables/{key}",
	Method:  "PUT",
}

var GetApiV4AdminClusters EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/admin/clusters",
	Method:  "GET",
}

var PostApiV4AdminClustersAdd EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/admin/clusters/add",
	Method:  "POST",
}

var DeleteApiV4AdminClustersByClusterId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/admin/clusters/{cluster_id}",
	Method:  "DELETE",
}

var GetApiV4AdminClustersByClusterId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/admin/clusters/{cluster_id}",
	Method:  "GET",
}

var PutApiV4AdminClustersByClusterId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/admin/clusters/{cluster_id}",
	Method:  "PUT",
}

var GetApiV4AdminDatabasesDictionaryTablesByDatabaseNameByTableName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/admin/databases/{database_name}/dictionary/tables/{table_name}",
	Method:  "GET",
}

var PostApiV4AdminMigrationsMarkByTimestamp EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/admin/migrations/{timestamp}/mark",
	Method:  "POST",
}

var GetApiV4ApplicationAppearance EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/application/appearance",
	Method:  "GET",
}

var PutApiV4ApplicationAppearance EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/application/appearance",
	Method:  "PUT",
}

var GetApiV4ApplicationPlanLimits EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/application/plan_limits",
	Method:  "GET",
}

var PutApiV4ApplicationPlanLimits EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/application/plan_limits",
	Method:  "PUT",
}

var GetApiV4ApplicationStatistics EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/application/statistics",
	Method:  "GET",
}

var GetApiV4Applications EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/applications",
	Method:  "GET",
}

var PostApiV4Applications EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/applications",
	Method:  "POST",
}

var DeleteApiV4ApplicationsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/applications/{id}",
	Method:  "DELETE",
}

var PostApiV4ApplicationsRenewSecretById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/applications/{id}/renew-secret",
	Method:  "POST",
}

var GetApiV4Avatar EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/avatar",
	Method:  "GET",
}

var GetApiV4BroadcastMessages EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/broadcast_messages",
	Method:  "GET",
}

var PostApiV4BroadcastMessages EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/broadcast_messages",
	Method:  "POST",
}

var DeleteApiV4BroadcastMessagesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/broadcast_messages/{id}",
	Method:  "DELETE",
}

var GetApiV4BroadcastMessagesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/broadcast_messages/{id}",
	Method:  "GET",
}

var PutApiV4BroadcastMessagesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/broadcast_messages/{id}",
	Method:  "PUT",
}

var GetApiV4BulkImports EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/bulk_imports",
	Method:  "GET",
}

var PostApiV4BulkImports EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/bulk_imports",
	Method:  "POST",
}

var GetApiV4BulkImportsEntities EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/bulk_imports/entities",
	Method:  "GET",
}

var GetApiV4BulkImportsByImportId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/bulk_imports/{import_id}",
	Method:  "GET",
}

var PostApiV4BulkImportsCancelByImportId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/bulk_imports/{import_id}/cancel",
	Method:  "POST",
}

var GetApiV4BulkImportsEntitiesByImportId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/bulk_imports/{import_id}/entities",
	Method:  "GET",
}

var GetApiV4BulkImportsEntitiesByImportIdByEntityId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/bulk_imports/{import_id}/entities/{entity_id}",
	Method:  "GET",
}

var GetApiV4BulkImportsEntitiesFailuresByImportIdByEntityId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/bulk_imports/{import_id}/entities/{entity_id}/failures",
	Method:  "GET",
}

var PostApiV4ContainerRegistryEventEvents EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/container_registry_event/events",
	Method:  "POST",
}

var GetApiV4DeployKeys EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/deploy_keys",
	Method:  "GET",
}

var PostApiV4DeployKeys EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/deploy_keys",
	Method:  "POST",
}

var GetApiV4DeployTokens EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/deploy_tokens",
	Method:  "GET",
}

var GetApiV4DiscoverCertBasedClusters EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/discover-cert-based-clusters",
	Method:  "GET",
}

var GetApiV4Events EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/events",
	Method:  "GET",
}

var GetApiV4FeatureFlagsUnleashByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/feature_flags/unleash/{project_id}",
	Method:  "GET",
}

var GetApiV4FeatureFlagsUnleashClientFeaturesByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/feature_flags/unleash/{project_id}/client/features",
	Method:  "GET",
}

var PostApiV4FeatureFlagsUnleashClientMetricsByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/feature_flags/unleash/{project_id}/client/metrics",
	Method:  "POST",
}

var PostApiV4FeatureFlagsUnleashClientRegisterByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/feature_flags/unleash/{project_id}/client/register",
	Method:  "POST",
}

var GetApiV4FeatureFlagsUnleashFeaturesByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/feature_flags/unleash/{project_id}/features",
	Method:  "GET",
}

var GetApiV4Features EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/features",
	Method:  "GET",
}

var GetApiV4FeaturesDefinitions EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/features/definitions",
	Method:  "GET",
}

var DeleteApiV4FeaturesByName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/features/{name}",
	Method:  "DELETE",
}

var PostApiV4FeaturesByName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/features/{name}",
	Method:  "POST",
}

var PostApiV4GeoNodeProxyGraphqlById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/geo/node_proxy/{id}/graphql",
	Method:  "POST",
}

var GetApiV4GeoProxy EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/geo/proxy",
	Method:  "GET",
}

var PostApiV4GeoProxyGitSshInfoRefsReceivePack EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/geo/proxy_git_ssh/info_refs_receive_pack",
	Method:  "POST",
}

var PostApiV4GeoProxyGitSshInfoRefsUploadPack EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/geo/proxy_git_ssh/info_refs_upload_pack",
	Method:  "POST",
}

var PostApiV4GeoProxyGitSshReceivePack EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/geo/proxy_git_ssh/receive_pack",
	Method:  "POST",
}

var PostApiV4GeoProxyGitSshUploadPack EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/geo/proxy_git_ssh/upload_pack",
	Method:  "POST",
}

var GetApiV4GeoRepositoriesPipelineRefsByGlRepository EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/geo/repositories/{gl_repository}/pipeline_refs",
	Method:  "GET",
}

var GetApiV4GeoRetrieveByReplicableNameByReplicableId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/geo/retrieve/{replicable_name}/{replicable_id}",
	Method:  "GET",
}

var PostApiV4GeoStatus EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/geo/status",
	Method:  "POST",
}

var GetApiV4GroupPackagesComposerByIdByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/group/{id}/-/packages/composer/{package_name:.+}",
	Method:  "GET",
}

var GetApiV4GroupPackagesComposerPByIdBySha EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/group/{id}/-/packages/composer/p/{sha}",
	Method:  "GET",
}

var GetApiV4GroupPackagesComposerP2ByIdByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/group/{id}/-/packages/composer/p2/{package_name:.+}",
	Method:  "GET",
}

var GetApiV4GroupPackagesComposerPackagesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/group/{id}/-/packages/composer/packages",
	Method:  "GET",
}

var GetApiV4Groups EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups",
	Method:  "GET",
}

var PostApiV4Groups EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups",
	Method:  "POST",
}

var PostApiV4GroupsImport EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/import",
	Method:  "POST",
}

var PostApiV4GroupsImportAuthorize EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/import/authorize",
	Method:  "POST",
}

var DeleteApiV4GroupsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}",
	Method:  "DELETE",
}

var GetApiV4GroupsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}",
	Method:  "GET",
}

var PutApiV4GroupsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}",
	Method:  "PUT",
}

var GetApiV4GroupsDebianDistributionsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/debian_distributions",
	Method:  "GET",
}

var PostApiV4GroupsDebianDistributionsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/debian_distributions",
	Method:  "POST",
}

var DeleteApiV4GroupsDebianDistributionsByIdByCodename EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/debian_distributions/{codename}",
	Method:  "DELETE",
}

var GetApiV4GroupsDebianDistributionsByIdByCodename EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/debian_distributions/{codename}",
	Method:  "GET",
}

var PutApiV4GroupsDebianDistributionsByIdByCodename EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/debian_distributions/{codename}",
	Method:  "PUT",
}

var GetApiV4GroupsDebianDistributionsKeyAscByIdByCodename EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/debian_distributions/{codename}/key.asc",
	Method:  "GET",
}

var GetApiV4GroupsPackagesDebianDistsInReleaseByIdByDistribution EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/debian/dists/{distribution:.+}/InRelease",
	Method:  "GET",
}

var GetApiV4GroupsPackagesDebianDistsReleaseByIdByDistribution EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/debian/dists/{distribution:.+}/Release",
	Method:  "GET",
}

var GetApiV4GroupsPackagesDebianDistsReleaseGpgByIdByDistribution EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/debian/dists/{distribution:.+}/Release.gpg",
	Method:  "GET",
}

var GetApiV4GroupsPackagesDebianDistsBinaryPackagesByIdByDistributionByComponentByArchitecture EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/debian/dists/{distribution:.+}/{component}/binary-{architecture}/Packages",
	Method:  "GET",
}

var GetApiV4GroupsPackagesDebianDistsBinaryByHashSHA256ByIdByDistributionByComponentByArchitectureByFileSha256 EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/debian/dists/{distribution:.+}/{component}/binary-{architecture}/by-hash/SHA256/{file_sha256}",
	Method:  "GET",
}

var GetApiV4GroupsPackagesDebianDistsDebianInstallerBinaryPackagesByIdByDistributionByComponentByArchitecture EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/debian/dists/{distribution:.+}/{component}/debian-installer/binary-{architecture}/Packages",
	Method:  "GET",
}

var GetApiV4GroupsPackagesDebianDistsDebianInstallerBinaryByHashSHA256ByIdByDistributionByComponentByArchitectureByFileSha256 EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/debian/dists/{distribution:.+}/{component}/debian-installer/binary-{architecture}/by-hash/SHA256/{file_sha256}",
	Method:  "GET",
}

var GetApiV4GroupsPackagesDebianDistsSourceSourcesByIdByDistributionByComponent EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/debian/dists/{distribution:.+}/{component}/source/Sources",
	Method:  "GET",
}

var GetApiV4GroupsPackagesDebianDistsSourceByHashSHA256ByIdByDistributionByComponentByFileSha256 EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/debian/dists/{distribution:.+}/{component}/source/by-hash/SHA256/{file_sha256}",
	Method:  "GET",
}

var GetApiV4GroupsPackagesDebianPoolByIdByDistributionByProjectIdByLetterByPackageNameByPackageVersionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/debian/pool/{distribution}/{project_id}/{letter}/{package_name}/{package_version}/{file_name}",
	Method:  "GET",
}

var GetApiV4GroupsPackagesMavenByIdByPathByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/maven/{path:.+}/{file_name}",
	Method:  "GET",
}

var GetApiV4GroupsPackagesNpmByIdByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/npm/{package_name:.+}",
	Method:  "GET",
}

var PostApiV4GroupsPackagesNpmNpmV1SecurityAdvisoriesBulkById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/npm/-/npm/v1/security/advisories/bulk",
	Method:  "POST",
}

var PostApiV4GroupsPackagesNpmNpmV1SecurityAuditsQuickById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/npm/-/npm/v1/security/audits/quick",
	Method:  "POST",
}

var GetApiV4GroupsPackagesNpmPackageDistTagsByIdByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/npm/-/package/{package_name:.+}/dist-tags",
	Method:  "GET",
}

var DeleteApiV4GroupsPackagesNpmPackageDistTagsByIdByPackageNameByTag EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/npm/-/package/{package_name:.+}/dist-tags/{tag}",
	Method:  "DELETE",
}

var PutApiV4GroupsPackagesNpmPackageDistTagsByIdByPackageNameByTag EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/npm/-/package/{package_name:.+}/dist-tags/{tag}",
	Method:  "PUT",
}

var GetApiV4GroupsPackagesNugetIndexById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/nuget/index",
	Method:  "GET",
}

var GetApiV4GroupsPackagesNugetMetadataByIdByPackageNameByPackageVersion EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/nuget/metadata/{package_name:.+}/{package_version:.+}",
	Method:  "GET",
}

var GetApiV4GroupsPackagesNugetMetadataIndexByIdByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/nuget/metadata/{package_name:.+}/index",
	Method:  "GET",
}

var GetApiV4GroupsPackagesNugetQueryById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/nuget/query",
	Method:  "GET",
}

var GetApiV4GroupsPackagesNugetSymbolfilesFileNameSignatureSameFileNameById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/nuget/symbolfiles/*file_name/*signature/*same_file_name",
	Method:  "GET",
}

var GetApiV4GroupsPackagesNugetV2ById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/nuget/v2",
	Method:  "GET",
}

var GetApiV4GroupsPackagesNugetV2MetadataById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/nuget/v2/$metadata",
	Method:  "GET",
}

var GetApiV4GroupsPackagesPypiFilesByIdBySha256ByFileIdentifier EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/pypi/files/{sha256}/{file_identifier:.+}",
	Method:  "GET",
}

var GetApiV4GroupsPackagesPypiSimpleById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/pypi/simple",
	Method:  "GET",
}

var GetApiV4GroupsPackagesPypiSimpleByIdByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/-/packages/pypi/simple/{package_name:.+}",
	Method:  "GET",
}

var GetApiV4GroupsAccessRequestsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/access_requests",
	Method:  "GET",
}

var PostApiV4GroupsAccessRequestsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/access_requests",
	Method:  "POST",
}

var DeleteApiV4GroupsAccessRequestsByIdByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/access_requests/{user_id}",
	Method:  "DELETE",
}

var PutApiV4GroupsAccessRequestsApproveByIdByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/access_requests/{user_id}/approve",
	Method:  "PUT",
}

var GetApiV4GroupsAccessTokensById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/access_tokens",
	Method:  "GET",
}

var PostApiV4GroupsAccessTokensById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/access_tokens",
	Method:  "POST",
}

var DeleteApiV4GroupsAccessTokensByIdByTokenId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/access_tokens/{token_id}",
	Method:  "DELETE",
}

var GetApiV4GroupsAccessTokensByIdByTokenId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/access_tokens/{token_id}",
	Method:  "GET",
}

var PostApiV4GroupsAccessTokensRotateByIdByTokenId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/access_tokens/{token_id}/rotate",
	Method:  "POST",
}

var PostApiV4GroupsAccessTokensSelfRotateById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/access_tokens/self/rotate",
	Method:  "POST",
}

var PostApiV4GroupsArchiveById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/archive",
	Method:  "POST",
}

var GetApiV4GroupsAuditEventsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/audit_events",
	Method:  "GET",
}

var GetApiV4GroupsAuditEventsByIdByAuditEventId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/audit_events/{audit_event_id}",
	Method:  "GET",
}

var GetApiV4GroupsAvatarById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/avatar",
	Method:  "GET",
}

var GetApiV4GroupsBadgesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/badges",
	Method:  "GET",
}

var PostApiV4GroupsBadgesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/badges",
	Method:  "POST",
}

var GetApiV4GroupsBadgesRenderById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/badges/render",
	Method:  "GET",
}

var DeleteApiV4GroupsBadgesByIdByBadgeId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/badges/{badge_id}",
	Method:  "DELETE",
}

var GetApiV4GroupsBadgesByIdByBadgeId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/badges/{badge_id}",
	Method:  "GET",
}

var PutApiV4GroupsBadgesByIdByBadgeId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/badges/{badge_id}",
	Method:  "PUT",
}

var GetApiV4GroupsBillableMembersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/billable_members",
	Method:  "GET",
}

var DeleteApiV4GroupsBillableMembersByIdByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/billable_members/{user_id}",
	Method:  "DELETE",
}

var GetApiV4GroupsBillableMembersIndirectByIdByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/billable_members/{user_id}/indirect",
	Method:  "GET",
}

var GetApiV4GroupsBillableMembersMembershipsByIdByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/billable_members/{user_id}/memberships",
	Method:  "GET",
}

var GetApiV4GroupsClustersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/clusters",
	Method:  "GET",
}

var PostApiV4GroupsClustersUserById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/clusters/user",
	Method:  "POST",
}

var DeleteApiV4GroupsClustersByIdByClusterId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/clusters/{cluster_id}",
	Method:  "DELETE",
}

var GetApiV4GroupsClustersByIdByClusterId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/clusters/{cluster_id}",
	Method:  "GET",
}

var PutApiV4GroupsClustersByIdByClusterId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/clusters/{cluster_id}",
	Method:  "PUT",
}

var GetApiV4GroupsCustomAttributesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/custom_attributes",
	Method:  "GET",
}

var DeleteApiV4GroupsCustomAttributesByIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/custom_attributes/{key}",
	Method:  "DELETE",
}

var GetApiV4GroupsCustomAttributesByIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/custom_attributes/{key}",
	Method:  "GET",
}

var PutApiV4GroupsCustomAttributesByIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/custom_attributes/{key}",
	Method:  "PUT",
}

var DeleteApiV4GroupsDependencyProxyCacheById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/dependency_proxy/cache",
	Method:  "DELETE",
}

var GetApiV4GroupsDeployTokensById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/deploy_tokens",
	Method:  "GET",
}

var PostApiV4GroupsDeployTokensById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/deploy_tokens",
	Method:  "POST",
}

var DeleteApiV4GroupsDeployTokensByIdByTokenId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/deploy_tokens/{token_id}",
	Method:  "DELETE",
}

var GetApiV4GroupsDeployTokensByIdByTokenId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/deploy_tokens/{token_id}",
	Method:  "GET",
}

var GetApiV4GroupsDescendantGroupsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/descendant_groups",
	Method:  "GET",
}

var GetApiV4GroupsEpicsAwardEmojiByIdByEpicIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/epics/{epic_iid}/award_emoji",
	Method:  "GET",
}

var PostApiV4GroupsEpicsAwardEmojiByIdByEpicIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/epics/{epic_iid}/award_emoji",
	Method:  "POST",
}

var DeleteApiV4GroupsEpicsAwardEmojiByIdByEpicIidByAwardId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/epics/{epic_iid}/award_emoji/{award_id}",
	Method:  "DELETE",
}

var GetApiV4GroupsEpicsAwardEmojiByIdByEpicIidByAwardId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/epics/{epic_iid}/award_emoji/{award_id}",
	Method:  "GET",
}

var GetApiV4GroupsEpicsNotesAwardEmojiByIdByEpicIidByNoteId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/epics/{epic_iid}/notes/{note_id}/award_emoji",
	Method:  "GET",
}

var PostApiV4GroupsEpicsNotesAwardEmojiByIdByEpicIidByNoteId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/epics/{epic_iid}/notes/{note_id}/award_emoji",
	Method:  "POST",
}

var DeleteApiV4GroupsEpicsNotesAwardEmojiByIdByEpicIidByNoteIdByAwardId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/epics/{epic_iid}/notes/{note_id}/award_emoji/{award_id}",
	Method:  "DELETE",
}

var GetApiV4GroupsEpicsNotesAwardEmojiByIdByEpicIidByNoteIdByAwardId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/epics/{epic_iid}/notes/{note_id}/award_emoji/{award_id}",
	Method:  "GET",
}

var PostApiV4GroupsExportById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/export",
	Method:  "POST",
}

var GetApiV4GroupsExportDownloadById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/export/download",
	Method:  "GET",
}

var PostApiV4GroupsExportRelationsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/export_relations",
	Method:  "POST",
}

var GetApiV4GroupsExportRelationsDownloadById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/export_relations/download",
	Method:  "GET",
}

var GetApiV4GroupsExportRelationsStatusById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/export_relations/status",
	Method:  "GET",
}

var GetApiV4GroupsGroupsSharedById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/groups/shared",
	Method:  "GET",
}

var GetApiV4GroupsIntegrationsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations",
	Method:  "GET",
}

var PutApiV4GroupsIntegrationsAppleAppStoreById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/apple-app-store",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsAsanaById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/asana",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsAssemblaById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/assembla",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsBambooById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/bamboo",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsBugzillaById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/bugzilla",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsBuildkiteById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/buildkite",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsCampfireById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/campfire",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsClickupById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/clickup",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsConfluenceById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/confluence",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsCustomIssueTrackerById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/custom-issue-tracker",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsDatadogById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/datadog",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsDiffblueCoverById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/diffblue-cover",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsDiscordById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/discord",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsDroneCiById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/drone-ci",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsEmailsOnPushById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/emails-on-push",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsEwmById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/ewm",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsExternalWikiById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/external-wiki",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsGitGuardianById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/git-guardian",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsGithubById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/github",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsGitlabSlackApplicationById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/gitlab-slack-application",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsGoogleCloudPlatformArtifactRegistryById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/google-cloud-platform-artifact-registry",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsGoogleCloudPlatformWorkloadIdentityFederationById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/google-cloud-platform-workload-identity-federation",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsGooglePlayById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/google-play",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsHangoutsChatById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/hangouts-chat",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsHarborById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/harbor",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsIrkerById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/irker",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsJenkinsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/jenkins",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsJiraById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/jira",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsJiraCloudAppById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/jira-cloud-app",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsLinearById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/linear",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsMatrixById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/matrix",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsMattermostById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/mattermost",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsMattermostSlashCommandsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/mattermost-slash-commands",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsMicrosoftTeamsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/microsoft-teams",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsMockCiById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/mock-ci",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsMockMonitoringById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/mock-monitoring",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsPackagistById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/packagist",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsPhorgeById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/phorge",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsPipelinesEmailById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/pipelines-email",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsPivotaltrackerById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/pivotaltracker",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsPumbleById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/pumble",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsPushoverById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/pushover",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsRedmineById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/redmine",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsSlackById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/slack",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsSlackSlashCommandsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/slack-slash-commands",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsSquashTmById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/squash-tm",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsTeamcityById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/teamcity",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsTelegramById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/telegram",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsUnifyCircuitById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/unify-circuit",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsWebexTeamsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/webex-teams",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsYoutrackById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/youtrack",
	Method:  "PUT",
}

var PutApiV4GroupsIntegrationsZentaoById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/zentao",
	Method:  "PUT",
}

var DeleteApiV4GroupsIntegrationsByIdBySlug EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/{slug}",
	Method:  "DELETE",
}

var GetApiV4GroupsIntegrationsByIdBySlug EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/integrations/{slug}",
	Method:  "GET",
}

var GetApiV4GroupsInvitationsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/invitations",
	Method:  "GET",
}

var PostApiV4GroupsInvitationsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/invitations",
	Method:  "POST",
}

var DeleteApiV4GroupsInvitationsByIdByEmail EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/invitations/{email}",
	Method:  "DELETE",
}

var PutApiV4GroupsInvitationsByIdByEmail EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/invitations/{email}",
	Method:  "PUT",
}

var GetApiV4GroupsInvitedGroupsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/invited_groups",
	Method:  "GET",
}

var GetApiV4GroupsIssuesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/issues",
	Method:  "GET",
}

var GetApiV4GroupsIssuesStatisticsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/issues_statistics",
	Method:  "GET",
}

var PostApiV4GroupsLdapSyncById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/ldap_sync",
	Method:  "POST",
}

var GetApiV4GroupsMembersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/members",
	Method:  "GET",
}

var PostApiV4GroupsMembersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/members",
	Method:  "POST",
}

var GetApiV4GroupsMembersAllById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/members/all",
	Method:  "GET",
}

var GetApiV4GroupsMembersAllByIdByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/members/all/{user_id}",
	Method:  "GET",
}

var PostApiV4GroupsMembersApproveAllById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/members/approve_all",
	Method:  "POST",
}

var PutApiV4GroupsMembersApproveByIdByMemberId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/members/{member_id}/approve",
	Method:  "PUT",
}

var DeleteApiV4GroupsMembersByIdByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/members/{user_id}",
	Method:  "DELETE",
}

var GetApiV4GroupsMembersByIdByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/members/{user_id}",
	Method:  "GET",
}

var PutApiV4GroupsMembersByIdByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/members/{user_id}",
	Method:  "PUT",
}

var DeleteApiV4GroupsMembersOverrideByIdByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/members/{user_id}/override",
	Method:  "DELETE",
}

var PostApiV4GroupsMembersOverrideByIdByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/members/{user_id}/override",
	Method:  "POST",
}

var PutApiV4GroupsMembersStateByIdByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/members/{user_id}/state",
	Method:  "PUT",
}

var GetApiV4GroupsMergeRequestsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/merge_requests",
	Method:  "GET",
}

var GetApiV4GroupsPackagesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/packages",
	Method:  "GET",
}

var GetApiV4GroupsPendingMembersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/pending_members",
	Method:  "GET",
}

var GetApiV4GroupsPlaceholderReassignmentsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/placeholder_reassignments",
	Method:  "GET",
}

var PostApiV4GroupsPlaceholderReassignmentsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/placeholder_reassignments",
	Method:  "POST",
}

var PostApiV4GroupsPlaceholderReassignmentsAuthorizeById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/placeholder_reassignments/authorize",
	Method:  "POST",
}

var GetApiV4GroupsProjectsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/projects",
	Method:  "GET",
}

var GetApiV4GroupsProjectsSharedById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/projects/shared",
	Method:  "GET",
}

var PostApiV4GroupsProjectsByIdByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/projects/{project_id}",
	Method:  "POST",
}

var GetApiV4GroupsProvisionedUsersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/provisioned_users",
	Method:  "GET",
}

var GetApiV4GroupsRegistryRepositoriesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/registry/repositories",
	Method:  "GET",
}

var GetApiV4GroupsReleasesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/releases",
	Method:  "GET",
}

var PostApiV4GroupsRestoreById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/restore",
	Method:  "POST",
}

var GetApiV4GroupsRunnersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/runners",
	Method:  "GET",
}

var PostApiV4GroupsRunnersResetRegistrationTokenById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/runners/reset_registration_token",
	Method:  "POST",
}

var GetApiV4GroupsSamlUsersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/saml_users",
	Method:  "GET",
}

var PostApiV4GroupsShareById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/share",
	Method:  "POST",
}

var DeleteApiV4GroupsShareByIdByGroupId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/share/{group_id}",
	Method:  "DELETE",
}

var GetApiV4GroupsSshCertificatesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/ssh_certificates",
	Method:  "GET",
}

var PostApiV4GroupsSshCertificatesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/ssh_certificates",
	Method:  "POST",
}

var DeleteApiV4GroupsSshCertificatesByIdBySshCertificatesId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/ssh_certificates/{ssh_certificates_id}",
	Method:  "DELETE",
}

var GetApiV4GroupsSubgroupsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/subgroups",
	Method:  "GET",
}

var PostApiV4GroupsTokensRevokeById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/tokens/revoke",
	Method:  "POST",
}

var PostApiV4GroupsTransferById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/transfer",
	Method:  "POST",
}

var GetApiV4GroupsTransferLocationsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/transfer_locations",
	Method:  "GET",
}

var PostApiV4GroupsUnarchiveById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/unarchive",
	Method:  "POST",
}

var GetApiV4GroupsUploadsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/uploads",
	Method:  "GET",
}

var DeleteApiV4GroupsUploadsByIdBySecretByFilename EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/uploads/{secret}/{filename}",
	Method:  "DELETE",
}

var GetApiV4GroupsUploadsByIdBySecretByFilename EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/uploads/{secret}/{filename}",
	Method:  "GET",
}

var DeleteApiV4GroupsUploadsByIdByUploadId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/uploads/{upload_id}",
	Method:  "DELETE",
}

var GetApiV4GroupsUploadsByIdByUploadId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/uploads/{upload_id}",
	Method:  "GET",
}

var GetApiV4GroupsUsersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/users",
	Method:  "GET",
}

var GetApiV4GroupsVariablesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/variables",
	Method:  "GET",
}

var PostApiV4GroupsVariablesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/variables",
	Method:  "POST",
}

var DeleteApiV4GroupsVariablesByIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/variables/{key}",
	Method:  "DELETE",
}

var GetApiV4GroupsVariablesByIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/variables/{key}",
	Method:  "GET",
}

var PutApiV4GroupsVariablesByIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/variables/{key}",
	Method:  "PUT",
}

var GetApiV4GroupsWikisById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/wikis",
	Method:  "GET",
}

var PostApiV4GroupsWikisById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/wikis",
	Method:  "POST",
}

var PostApiV4GroupsWikisAttachmentsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/wikis/attachments",
	Method:  "POST",
}

var DeleteApiV4GroupsWikisByIdBySlug EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/wikis/{slug}",
	Method:  "DELETE",
}

var GetApiV4GroupsWikisByIdBySlug EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/wikis/{slug}",
	Method:  "GET",
}

var PutApiV4GroupsWikisByIdBySlug EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/groups/{id}/wikis/{slug}",
	Method:  "PUT",
}

var GetApiV4Hooks EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/hooks",
	Method:  "GET",
}

var PostApiV4Hooks EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/hooks",
	Method:  "POST",
}

var DeleteApiV4HooksByHookId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/hooks/{hook_id}",
	Method:  "DELETE",
}

var GetApiV4HooksByHookId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/hooks/{hook_id}",
	Method:  "GET",
}

var PostApiV4HooksByHookId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/hooks/{hook_id}",
	Method:  "POST",
}

var PutApiV4HooksByHookId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/hooks/{hook_id}",
	Method:  "PUT",
}

var DeleteApiV4HooksCustomHeadersByHookIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/hooks/{hook_id}/custom_headers/{key}",
	Method:  "DELETE",
}

var PutApiV4HooksCustomHeadersByHookIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/hooks/{hook_id}/custom_headers/{key}",
	Method:  "PUT",
}

var DeleteApiV4HooksUrlVariablesByHookIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/hooks/{hook_id}/url_variables/{key}",
	Method:  "DELETE",
}

var PutApiV4HooksUrlVariablesByHookIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/hooks/{hook_id}/url_variables/{key}",
	Method:  "PUT",
}

var PostApiV4ImportBitbucket EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/import/bitbucket",
	Method:  "POST",
}

var PostApiV4ImportBitbucketServer EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/import/bitbucket_server",
	Method:  "POST",
}

var PostApiV4ImportGithub EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/import/github",
	Method:  "POST",
}

var PostApiV4ImportGithubCancel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/import/github/cancel",
	Method:  "POST",
}

var PostApiV4ImportGithubGists EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/import/github/gists",
	Method:  "POST",
}

var PostApiV4IntegrationsJiraConnectSubscriptions EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/integrations/jira_connect/subscriptions",
	Method:  "POST",
}

var PostApiV4IntegrationsSlackEvents EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/integrations/slack/events",
	Method:  "POST",
}

var PostApiV4IntegrationsSlackInteractions EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/integrations/slack/interactions",
	Method:  "POST",
}

var PostApiV4IntegrationsSlackOptions EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/integrations/slack/options",
	Method:  "POST",
}

var GetApiV4Issues EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/issues",
	Method:  "GET",
}

var GetApiV4IssuesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/issues/{id}",
	Method:  "GET",
}

var GetApiV4IssuesStatistics EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/issues_statistics",
	Method:  "GET",
}

var GetApiV4Job EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/job",
	Method:  "GET",
}

var GetApiV4JobAllowedAgents EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/job/allowed_agents",
	Method:  "GET",
}

var PostApiV4JobsRequest EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/jobs/request",
	Method:  "POST",
}

var PutApiV4JobsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/jobs/{id}",
	Method:  "PUT",
}

var GetApiV4JobsArtifactsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/jobs/{id}/artifacts",
	Method:  "GET",
}

var PostApiV4JobsArtifactsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/jobs/{id}/artifacts",
	Method:  "POST",
}

var PostApiV4JobsArtifactsAuthorizeById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/jobs/{id}/artifacts/authorize",
	Method:  "POST",
}

var PatchApiV4JobsTraceById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/jobs/{id}/trace",
	Method:  "PATCH",
}

var GetApiV4Keys EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/keys",
	Method:  "GET",
}

var GetApiV4KeysById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/keys/{id}",
	Method:  "GET",
}

var PostApiV4Markdown EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/markdown",
	Method:  "POST",
}

var GetApiV4MergeRequests EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/merge_requests",
	Method:  "GET",
}

var GetApiV4Metadata EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/metadata",
	Method:  "GET",
}

var GetApiV4Namespaces EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/namespaces",
	Method:  "GET",
}

var GetApiV4NamespacesStorageLimitExclusions EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/namespaces/storage/limit_exclusions",
	Method:  "GET",
}

var GetApiV4NamespacesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/namespaces/{id}",
	Method:  "GET",
}

var PutApiV4NamespacesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/namespaces/{id}",
	Method:  "PUT",
}

var GetApiV4NamespacesExistsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/namespaces/{id}/exists",
	Method:  "GET",
}

var GetApiV4NamespacesGitlabSubscriptionById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/namespaces/{id}/gitlab_subscription",
	Method:  "GET",
}

var DeleteApiV4NamespacesStorageLimitExclusionById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/namespaces/{id}/storage/limit_exclusion",
	Method:  "DELETE",
}

var PostApiV4NamespacesStorageLimitExclusionById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/namespaces/{id}/storage/limit_exclusion",
	Method:  "POST",
}

var PostApiV4Organizations EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/organizations",
	Method:  "POST",
}

var GetApiV4PackagesConanV1ConansSearch EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/conans/search",
	Method:  "GET",
}

var DeleteApiV4PackagesConanV1ConansByPackageNameByPackageVersionByPackageUsernameByPackageChannel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}",
	Method:  "DELETE",
}

var GetApiV4PackagesConanV1ConansByPackageNameByPackageVersionByPackageUsernameByPackageChannel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}",
	Method:  "GET",
}

var GetApiV4PackagesConanV1ConansDigestByPackageNameByPackageVersionByPackageUsernameByPackageChannel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/digest",
	Method:  "GET",
}

var GetApiV4PackagesConanV1ConansDownloadUrlsByPackageNameByPackageVersionByPackageUsernameByPackageChannel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/download_urls",
	Method:  "GET",
}

var GetApiV4PackagesConanV1ConansPackagesByPackageNameByPackageVersionByPackageUsernameByPackageChannelByConanPackageReference EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/packages/{conan_package_reference}",
	Method:  "GET",
}

var GetApiV4PackagesConanV1ConansPackagesDigestByPackageNameByPackageVersionByPackageUsernameByPackageChannelByConanPackageReference EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/packages/{conan_package_reference}/digest",
	Method:  "GET",
}

var GetApiV4PackagesConanV1ConansPackagesDownloadUrlsByPackageNameByPackageVersionByPackageUsernameByPackageChannelByConanPackageReference EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/packages/{conan_package_reference}/download_urls",
	Method:  "GET",
}

var PostApiV4PackagesConanV1ConansPackagesUploadUrlsByPackageNameByPackageVersionByPackageUsernameByPackageChannelByConanPackageReference EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/packages/{conan_package_reference}/upload_urls",
	Method:  "POST",
}

var GetApiV4PackagesConanV1ConansSearchByPackageNameByPackageVersionByPackageUsernameByPackageChannel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/search",
	Method:  "GET",
}

var PostApiV4PackagesConanV1ConansUploadUrlsByPackageNameByPackageVersionByPackageUsernameByPackageChannel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/upload_urls",
	Method:  "POST",
}

var GetApiV4PackagesConanV1FilesExportByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/export/{file_name}",
	Method:  "GET",
}

var PutApiV4PackagesConanV1FilesExportByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/export/{file_name}",
	Method:  "PUT",
}

var PutApiV4PackagesConanV1FilesExportAuthorizeByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/export/{file_name}/authorize",
	Method:  "PUT",
}

var GetApiV4PackagesConanV1FilesPackageByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByConanPackageReferenceByPackageRevisionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/package/{conan_package_reference}/{package_revision}/{file_name}",
	Method:  "GET",
}

var PutApiV4PackagesConanV1FilesPackageByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByConanPackageReferenceByPackageRevisionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/package/{conan_package_reference}/{package_revision}/{file_name}",
	Method:  "PUT",
}

var PutApiV4PackagesConanV1FilesPackageAuthorizeByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByConanPackageReferenceByPackageRevisionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/package/{conan_package_reference}/{package_revision}/{file_name}/authorize",
	Method:  "PUT",
}

var GetApiV4PackagesConanV1Ping EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/ping",
	Method:  "GET",
}

var GetApiV4PackagesConanV1UsersAuthenticate EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/users/authenticate",
	Method:  "GET",
}

var GetApiV4PackagesConanV1UsersCheckCredentials EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/conan/v1/users/check_credentials",
	Method:  "GET",
}

var GetApiV4PackagesMavenByPathByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/maven/{path:.+}/{file_name}",
	Method:  "GET",
}

var GetApiV4PackagesNpmByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/npm/{package_name:.+}",
	Method:  "GET",
}

var PostApiV4PackagesNpmNpmV1SecurityAdvisoriesBulk EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/npm/-/npm/v1/security/advisories/bulk",
	Method:  "POST",
}

var PostApiV4PackagesNpmNpmV1SecurityAuditsQuick EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/npm/-/npm/v1/security/audits/quick",
	Method:  "POST",
}

var GetApiV4PackagesNpmPackageDistTagsByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/npm/-/package/{package_name:.+}/dist-tags",
	Method:  "GET",
}

var DeleteApiV4PackagesNpmPackageDistTagsByPackageNameByTag EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/npm/-/package/{package_name:.+}/dist-tags/{tag}",
	Method:  "DELETE",
}

var PutApiV4PackagesNpmPackageDistTagsByPackageNameByTag EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/npm/-/package/{package_name:.+}/dist-tags/{tag}",
	Method:  "PUT",
}

var GetApiV4PackagesTerraformModulesV1ByModuleNamespaceByModuleNameByModuleSystem EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/terraform/modules/v1/{module_namespace}/{module_name}/{module_system}",
	Method:  "GET",
}

var GetApiV4PackagesTerraformModulesV1ByModuleNamespaceByModuleNameByModuleSystemByModuleVersion EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/terraform/modules/v1/{module_namespace}/{module_name}/{module_system}/{module_version:.+}",
	Method:  "GET",
}

var GetApiV4PackagesTerraformModulesV1DownloadByModuleNamespaceByModuleNameByModuleSystemByModuleVersion EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/terraform/modules/v1/{module_namespace}/{module_name}/{module_system}/{module_version:.+}/download",
	Method:  "GET",
}

var GetApiV4PackagesTerraformModulesV1FileByModuleNamespaceByModuleNameByModuleSystemByModuleVersion EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/terraform/modules/v1/{module_namespace}/{module_name}/{module_system}/{module_version:.+}/file",
	Method:  "GET",
}

var GetApiV4PackagesTerraformModulesV1DownloadByModuleNamespaceByModuleNameByModuleSystem EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/terraform/modules/v1/{module_namespace}/{module_name}/{module_system}/download",
	Method:  "GET",
}

var GetApiV4PackagesTerraformModulesV1VersionsByModuleNamespaceByModuleNameByModuleSystem EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/packages/terraform/modules/v1/{module_namespace}/{module_name}/{module_system}/versions",
	Method:  "GET",
}

var GetApiV4PagesDomains EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/pages/domains",
	Method:  "GET",
}

var GetApiV4PersonalAccessTokens EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/personal_access_tokens",
	Method:  "GET",
}

var DeleteApiV4PersonalAccessTokensSelf EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/personal_access_tokens/self",
	Method:  "DELETE",
}

var GetApiV4PersonalAccessTokensSelf EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/personal_access_tokens/self",
	Method:  "GET",
}

var GetApiV4PersonalAccessTokensSelfAssociations EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/personal_access_tokens/self/associations",
	Method:  "GET",
}

var PostApiV4PersonalAccessTokensSelfRotate EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/personal_access_tokens/self/rotate",
	Method:  "POST",
}

var DeleteApiV4PersonalAccessTokensById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/personal_access_tokens/{id}",
	Method:  "DELETE",
}

var GetApiV4PersonalAccessTokensById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/personal_access_tokens/{id}",
	Method:  "GET",
}

var PostApiV4PersonalAccessTokensRotateById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/personal_access_tokens/{id}/rotate",
	Method:  "POST",
}

var GetApiV4Projects EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects",
	Method:  "GET",
}

var PostApiV4Projects EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects",
	Method:  "POST",
}

var PostApiV4ProjectsImport EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/import",
	Method:  "POST",
}

var PostApiV4ProjectsImportRelation EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/import-relation",
	Method:  "POST",
}

var PostApiV4ProjectsImportRelationAuthorize EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/import-relation/authorize",
	Method:  "POST",
}

var PostApiV4ProjectsImportAuthorize EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/import/authorize",
	Method:  "POST",
}

var PostApiV4ProjectsRemoteImport EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/remote-import",
	Method:  "POST",
}

var PostApiV4ProjectsRemoteImportS3 EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/remote-import-s3",
	Method:  "POST",
}

var PostApiV4ProjectsUserByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/user/{user_id}",
	Method:  "POST",
}

var DeleteApiV4ProjectsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}",
	Method:  "GET",
}

var PutApiV4ProjectsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}",
	Method:  "PUT",
}

var PostApiV4ProjectsRefTriggerPipelineByIdByRef EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/(ref/{ref}/)trigger/pipeline",
	Method:  "POST",
}

var GetApiV4ProjectsAccessRequestsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/access_requests",
	Method:  "GET",
}

var PostApiV4ProjectsAccessRequestsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/access_requests",
	Method:  "POST",
}

var DeleteApiV4ProjectsAccessRequestsByIdByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/access_requests/{user_id}",
	Method:  "DELETE",
}

var PutApiV4ProjectsAccessRequestsApproveByIdByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/access_requests/{user_id}/approve",
	Method:  "PUT",
}

var GetApiV4ProjectsAccessTokensById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/access_tokens",
	Method:  "GET",
}

var PostApiV4ProjectsAccessTokensById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/access_tokens",
	Method:  "POST",
}

var DeleteApiV4ProjectsAccessTokensByIdByTokenId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/access_tokens/{token_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsAccessTokensByIdByTokenId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/access_tokens/{token_id}",
	Method:  "GET",
}

var PostApiV4ProjectsAccessTokensRotateByIdByTokenId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/access_tokens/{token_id}/rotate",
	Method:  "POST",
}

var PostApiV4ProjectsAccessTokensSelfRotateById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/access_tokens/self/rotate",
	Method:  "POST",
}

var GetApiV4ProjectsAlertManagementAlertsMetricImagesByIdByAlertIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/alert_management_alerts/{alert_iid}/metric_images",
	Method:  "GET",
}

var PostApiV4ProjectsAlertManagementAlertsMetricImagesByIdByAlertIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/alert_management_alerts/{alert_iid}/metric_images",
	Method:  "POST",
}

var PostApiV4ProjectsAlertManagementAlertsMetricImagesAuthorizeByIdByAlertIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/alert_management_alerts/{alert_iid}/metric_images/authorize",
	Method:  "POST",
}

var DeleteApiV4ProjectsAlertManagementAlertsMetricImagesByIdByAlertIidByMetricImageId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/alert_management_alerts/{alert_iid}/metric_images/{metric_image_id}",
	Method:  "DELETE",
}

var PutApiV4ProjectsAlertManagementAlertsMetricImagesByIdByAlertIidByMetricImageId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/alert_management_alerts/{alert_iid}/metric_images/{metric_image_id}",
	Method:  "PUT",
}

var PostApiV4ProjectsArchiveById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/archive",
	Method:  "POST",
}

var DeleteApiV4ProjectsArtifactsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/artifacts",
	Method:  "DELETE",
}

var GetApiV4ProjectsAuditEventsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/audit_events",
	Method:  "GET",
}

var GetApiV4ProjectsAuditEventsByIdByAuditEventId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/audit_events/{audit_event_id}",
	Method:  "GET",
}

var GetApiV4ProjectsAvatarById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/avatar",
	Method:  "GET",
}

var GetApiV4ProjectsBadgesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/badges",
	Method:  "GET",
}

var PostApiV4ProjectsBadgesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/badges",
	Method:  "POST",
}

var GetApiV4ProjectsBadgesRenderById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/badges/render",
	Method:  "GET",
}

var DeleteApiV4ProjectsBadgesByIdByBadgeId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/badges/{badge_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsBadgesByIdByBadgeId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/badges/{badge_id}",
	Method:  "GET",
}

var PutApiV4ProjectsBadgesByIdByBadgeId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/badges/{badge_id}",
	Method:  "PUT",
}

var PostApiV4ProjectsCatalogPublishById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/catalog/publish",
	Method:  "POST",
}

var GetApiV4ProjectsCiLintById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/ci/lint",
	Method:  "GET",
}

var PostApiV4ProjectsCiLintById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/ci/lint",
	Method:  "POST",
}

var GetApiV4ProjectsClusterAgentsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/cluster_agents",
	Method:  "GET",
}

var PostApiV4ProjectsClusterAgentsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/cluster_agents",
	Method:  "POST",
}

var DeleteApiV4ProjectsClusterAgentsByIdByAgentId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/cluster_agents/{agent_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsClusterAgentsByIdByAgentId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/cluster_agents/{agent_id}",
	Method:  "GET",
}

var GetApiV4ProjectsClusterAgentsTokensByIdByAgentId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/cluster_agents/{agent_id}/tokens",
	Method:  "GET",
}

var PostApiV4ProjectsClusterAgentsTokensByIdByAgentId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/cluster_agents/{agent_id}/tokens",
	Method:  "POST",
}

var DeleteApiV4ProjectsClusterAgentsTokensByIdByAgentIdByTokenId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/cluster_agents/{agent_id}/tokens/{token_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsClusterAgentsTokensByIdByAgentIdByTokenId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/cluster_agents/{agent_id}/tokens/{token_id}",
	Method:  "GET",
}

var GetApiV4ProjectsClustersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/clusters",
	Method:  "GET",
}

var PostApiV4ProjectsClustersUserById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/clusters/user",
	Method:  "POST",
}

var DeleteApiV4ProjectsClustersByIdByClusterId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/clusters/{cluster_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsClustersByIdByClusterId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/clusters/{cluster_id}",
	Method:  "GET",
}

var PutApiV4ProjectsClustersByIdByClusterId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/clusters/{cluster_id}",
	Method:  "PUT",
}

var PostApiV4ProjectsCreateCiConfigById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/create_ci_config",
	Method:  "POST",
}

var GetApiV4ProjectsCustomAttributesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/custom_attributes",
	Method:  "GET",
}

var DeleteApiV4ProjectsCustomAttributesByIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/custom_attributes/{key}",
	Method:  "DELETE",
}

var GetApiV4ProjectsCustomAttributesByIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/custom_attributes/{key}",
	Method:  "GET",
}

var PutApiV4ProjectsCustomAttributesByIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/custom_attributes/{key}",
	Method:  "PUT",
}

var GetApiV4ProjectsDebianDistributionsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/debian_distributions",
	Method:  "GET",
}

var PostApiV4ProjectsDebianDistributionsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/debian_distributions",
	Method:  "POST",
}

var DeleteApiV4ProjectsDebianDistributionsByIdByCodename EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/debian_distributions/{codename}",
	Method:  "DELETE",
}

var GetApiV4ProjectsDebianDistributionsByIdByCodename EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/debian_distributions/{codename}",
	Method:  "GET",
}

var PutApiV4ProjectsDebianDistributionsByIdByCodename EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/debian_distributions/{codename}",
	Method:  "PUT",
}

var GetApiV4ProjectsDebianDistributionsKeyAscByIdByCodename EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/debian_distributions/{codename}/key.asc",
	Method:  "GET",
}

var GetApiV4ProjectsDeployKeysById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/deploy_keys",
	Method:  "GET",
}

var PostApiV4ProjectsDeployKeysById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/deploy_keys",
	Method:  "POST",
}

var DeleteApiV4ProjectsDeployKeysByIdByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/deploy_keys/{key_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsDeployKeysByIdByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/deploy_keys/{key_id}",
	Method:  "GET",
}

var PutApiV4ProjectsDeployKeysByIdByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/deploy_keys/{key_id}",
	Method:  "PUT",
}

var PostApiV4ProjectsDeployKeysEnableByIdByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/deploy_keys/{key_id}/enable",
	Method:  "POST",
}

var GetApiV4ProjectsDeployTokensById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/deploy_tokens",
	Method:  "GET",
}

var PostApiV4ProjectsDeployTokensById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/deploy_tokens",
	Method:  "POST",
}

var DeleteApiV4ProjectsDeployTokensByIdByTokenId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/deploy_tokens/{token_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsDeployTokensByIdByTokenId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/deploy_tokens/{token_id}",
	Method:  "GET",
}

var GetApiV4ProjectsDeploymentsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/deployments",
	Method:  "GET",
}

var PostApiV4ProjectsDeploymentsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/deployments",
	Method:  "POST",
}

var DeleteApiV4ProjectsDeploymentsByIdByDeploymentId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/deployments/{deployment_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsDeploymentsByIdByDeploymentId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/deployments/{deployment_id}",
	Method:  "GET",
}

var PutApiV4ProjectsDeploymentsByIdByDeploymentId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/deployments/{deployment_id}",
	Method:  "PUT",
}

var PostApiV4ProjectsDeploymentsApprovalByIdByDeploymentId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/deployments/{deployment_id}/approval",
	Method:  "POST",
}

var GetApiV4ProjectsDeploymentsMergeRequestsByIdByDeploymentId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/deployments/{deployment_id}/merge_requests",
	Method:  "GET",
}

var GetApiV4ProjectsEnvironmentsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/environments",
	Method:  "GET",
}

var PostApiV4ProjectsEnvironmentsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/environments",
	Method:  "POST",
}

var DeleteApiV4ProjectsEnvironmentsReviewAppsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/environments/review_apps",
	Method:  "DELETE",
}

var PostApiV4ProjectsEnvironmentsStopStaleById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/environments/stop_stale",
	Method:  "POST",
}

var DeleteApiV4ProjectsEnvironmentsByIdByEnvironmentId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/environments/{environment_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsEnvironmentsByIdByEnvironmentId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/environments/{environment_id}",
	Method:  "GET",
}

var PutApiV4ProjectsEnvironmentsByIdByEnvironmentId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/environments/{environment_id}",
	Method:  "PUT",
}

var PostApiV4ProjectsEnvironmentsStopByIdByEnvironmentId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/environments/{environment_id}/stop",
	Method:  "POST",
}

var GetApiV4ProjectsErrorTrackingClientKeysById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/error_tracking/client_keys",
	Method:  "GET",
}

var PostApiV4ProjectsErrorTrackingClientKeysById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/error_tracking/client_keys",
	Method:  "POST",
}

var DeleteApiV4ProjectsErrorTrackingClientKeysByIdByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/error_tracking/client_keys/{key_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsErrorTrackingSettingsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/error_tracking/settings",
	Method:  "GET",
}

var PatchApiV4ProjectsErrorTrackingSettingsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/error_tracking/settings",
	Method:  "PATCH",
}

var PutApiV4ProjectsErrorTrackingSettingsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/error_tracking/settings",
	Method:  "PUT",
}

var GetApiV4ProjectsEventsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/events",
	Method:  "GET",
}

var GetApiV4ProjectsExportById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/export",
	Method:  "GET",
}

var PostApiV4ProjectsExportById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/export",
	Method:  "POST",
}

var GetApiV4ProjectsExportDownloadById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/export/download",
	Method:  "GET",
}

var PostApiV4ProjectsExportRelationsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/export_relations",
	Method:  "POST",
}

var GetApiV4ProjectsExportRelationsDownloadById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/export_relations/download",
	Method:  "GET",
}

var GetApiV4ProjectsExportRelationsStatusById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/export_relations/status",
	Method:  "GET",
}

var GetApiV4ProjectsFeatureFlagsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/feature_flags",
	Method:  "GET",
}

var PostApiV4ProjectsFeatureFlagsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/feature_flags",
	Method:  "POST",
}

var DeleteApiV4ProjectsFeatureFlagsByIdByFeatureFlagName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/feature_flags/{feature_flag_name}",
	Method:  "DELETE",
}

var GetApiV4ProjectsFeatureFlagsByIdByFeatureFlagName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/feature_flags/{feature_flag_name}",
	Method:  "GET",
}

var PutApiV4ProjectsFeatureFlagsByIdByFeatureFlagName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/feature_flags/{feature_flag_name}",
	Method:  "PUT",
}

var GetApiV4ProjectsFeatureFlagsUserListsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/feature_flags_user_lists",
	Method:  "GET",
}

var PostApiV4ProjectsFeatureFlagsUserListsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/feature_flags_user_lists",
	Method:  "POST",
}

var DeleteApiV4ProjectsFeatureFlagsUserListsByIdByIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/feature_flags_user_lists/{iid}",
	Method:  "DELETE",
}

var GetApiV4ProjectsFeatureFlagsUserListsByIdByIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/feature_flags_user_lists/{iid}",
	Method:  "GET",
}

var PutApiV4ProjectsFeatureFlagsUserListsByIdByIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/feature_flags_user_lists/{iid}",
	Method:  "PUT",
}

var DeleteApiV4ProjectsForkById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/fork",
	Method:  "DELETE",
}

var PostApiV4ProjectsForkById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/fork",
	Method:  "POST",
}

var PostApiV4ProjectsForkByIdByForkedFromId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/fork/{forked_from_id}",
	Method:  "POST",
}

var GetApiV4ProjectsForksById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/forks",
	Method:  "GET",
}

var GetApiV4ProjectsFreezePeriodsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/freeze_periods",
	Method:  "GET",
}

var PostApiV4ProjectsFreezePeriodsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/freeze_periods",
	Method:  "POST",
}

var DeleteApiV4ProjectsFreezePeriodsByIdByFreezePeriodId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/freeze_periods/{freeze_period_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsFreezePeriodsByIdByFreezePeriodId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/freeze_periods/{freeze_period_id}",
	Method:  "GET",
}

var PutApiV4ProjectsFreezePeriodsByIdByFreezePeriodId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/freeze_periods/{freeze_period_id}",
	Method:  "PUT",
}

var GetApiV4ProjectsGroupsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/groups",
	Method:  "GET",
}

var GetApiV4ProjectsHooksById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/hooks",
	Method:  "GET",
}

var PostApiV4ProjectsHooksById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/hooks",
	Method:  "POST",
}

var DeleteApiV4ProjectsHooksByIdByHookId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/hooks/{hook_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsHooksByIdByHookId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/hooks/{hook_id}",
	Method:  "GET",
}

var PutApiV4ProjectsHooksByIdByHookId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/hooks/{hook_id}",
	Method:  "PUT",
}

var DeleteApiV4ProjectsHooksCustomHeadersByIdByHookIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/hooks/{hook_id}/custom_headers/{key}",
	Method:  "DELETE",
}

var PutApiV4ProjectsHooksCustomHeadersByIdByHookIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/hooks/{hook_id}/custom_headers/{key}",
	Method:  "PUT",
}

var GetApiV4ProjectsHooksEventsByIdByHookId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/hooks/{hook_id}/events",
	Method:  "GET",
}

var PostApiV4ProjectsHooksEventsResendByIdByHookIdByHookLogId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/hooks/{hook_id}/events/{hook_log_id}/resend",
	Method:  "POST",
}

var PostApiV4ProjectsHooksTestByIdByHookIdByTrigger EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/hooks/{hook_id}/test/{trigger}",
	Method:  "POST",
}

var DeleteApiV4ProjectsHooksUrlVariablesByIdByHookIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/hooks/{hook_id}/url_variables/{key}",
	Method:  "DELETE",
}

var PutApiV4ProjectsHooksUrlVariablesByIdByHookIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/hooks/{hook_id}/url_variables/{key}",
	Method:  "PUT",
}

var PostApiV4ProjectsHousekeepingById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/housekeeping",
	Method:  "POST",
}

var GetApiV4ProjectsImportById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/import",
	Method:  "GET",
}

var PostApiV4ProjectsImportProjectMembersByIdByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/import_project_members/{project_id}",
	Method:  "POST",
}

var GetApiV4ProjectsIntegrationsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations",
	Method:  "GET",
}

var PutApiV4ProjectsIntegrationsAppleAppStoreById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/apple-app-store",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsAsanaById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/asana",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsAssemblaById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/assembla",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsBambooById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/bamboo",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsBugzillaById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/bugzilla",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsBuildkiteById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/buildkite",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsCampfireById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/campfire",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsClickupById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/clickup",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsConfluenceById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/confluence",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsCustomIssueTrackerById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/custom-issue-tracker",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsDatadogById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/datadog",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsDiffblueCoverById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/diffblue-cover",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsDiscordById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/discord",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsDroneCiById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/drone-ci",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsEmailsOnPushById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/emails-on-push",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsEwmById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/ewm",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsExternalWikiById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/external-wiki",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsGitGuardianById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/git-guardian",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsGithubById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/github",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsGitlabSlackApplicationById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/gitlab-slack-application",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsGoogleCloudPlatformArtifactRegistryById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/google-cloud-platform-artifact-registry",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsGoogleCloudPlatformWorkloadIdentityFederationById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/google-cloud-platform-workload-identity-federation",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsGooglePlayById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/google-play",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsHangoutsChatById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/hangouts-chat",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsHarborById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/harbor",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsIrkerById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/irker",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsJenkinsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/jenkins",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsJiraById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/jira",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsJiraCloudAppById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/jira-cloud-app",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsLinearById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/linear",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsMatrixById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/matrix",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsMattermostById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/mattermost",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsMattermostSlashCommandsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/mattermost-slash-commands",
	Method:  "PUT",
}

var PostApiV4ProjectsIntegrationsMattermostSlashCommandsTriggerById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/mattermost_slash_commands/trigger",
	Method:  "POST",
}

var PutApiV4ProjectsIntegrationsMicrosoftTeamsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/microsoft-teams",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsMockCiById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/mock-ci",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsMockMonitoringById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/mock-monitoring",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsPackagistById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/packagist",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsPhorgeById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/phorge",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsPipelinesEmailById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/pipelines-email",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsPivotaltrackerById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/pivotaltracker",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsPumbleById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/pumble",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsPushoverById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/pushover",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsRedmineById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/redmine",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsSlackById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/slack",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsSlackSlashCommandsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/slack-slash-commands",
	Method:  "PUT",
}

var PostApiV4ProjectsIntegrationsSlackSlashCommandsTriggerById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/slack_slash_commands/trigger",
	Method:  "POST",
}

var PutApiV4ProjectsIntegrationsSquashTmById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/squash-tm",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsTeamcityById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/teamcity",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsTelegramById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/telegram",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsUnifyCircuitById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/unify-circuit",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsWebexTeamsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/webex-teams",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsYoutrackById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/youtrack",
	Method:  "PUT",
}

var PutApiV4ProjectsIntegrationsZentaoById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/zentao",
	Method:  "PUT",
}

var DeleteApiV4ProjectsIntegrationsByIdBySlug EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/{slug}",
	Method:  "DELETE",
}

var GetApiV4ProjectsIntegrationsByIdBySlug EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/integrations/{slug}",
	Method:  "GET",
}

var GetApiV4ProjectsInvitationsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/invitations",
	Method:  "GET",
}

var PostApiV4ProjectsInvitationsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/invitations",
	Method:  "POST",
}

var DeleteApiV4ProjectsInvitationsByIdByEmail EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/invitations/{email}",
	Method:  "DELETE",
}

var PutApiV4ProjectsInvitationsByIdByEmail EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/invitations/{email}",
	Method:  "PUT",
}

var GetApiV4ProjectsInvitedGroupsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/invited_groups",
	Method:  "GET",
}

var GetApiV4ProjectsIssuesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues",
	Method:  "GET",
}

var PostApiV4ProjectsIssuesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues",
	Method:  "POST",
}

var GetApiV4ProjectsIssuesResourceMilestoneEventsByIdByEventableId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{eventable_id}/resource_milestone_events",
	Method:  "GET",
}

var GetApiV4ProjectsIssuesResourceMilestoneEventsByIdByEventableIdByEventId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{eventable_id}/resource_milestone_events/{event_id}",
	Method:  "GET",
}

var DeleteApiV4ProjectsIssuesByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}",
	Method:  "DELETE",
}

var GetApiV4ProjectsIssuesByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}",
	Method:  "GET",
}

var PutApiV4ProjectsIssuesByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}",
	Method:  "PUT",
}

var PostApiV4ProjectsIssuesAddSpentTimeByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/add_spent_time",
	Method:  "POST",
}

var GetApiV4ProjectsIssuesAwardEmojiByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/award_emoji",
	Method:  "GET",
}

var PostApiV4ProjectsIssuesAwardEmojiByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/award_emoji",
	Method:  "POST",
}

var DeleteApiV4ProjectsIssuesAwardEmojiByIdByIssueIidByAwardId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/award_emoji/{award_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsIssuesAwardEmojiByIdByIssueIidByAwardId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/award_emoji/{award_id}",
	Method:  "GET",
}

var PostApiV4ProjectsIssuesCloneByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/clone",
	Method:  "POST",
}

var GetApiV4ProjectsIssuesClosedByByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/closed_by",
	Method:  "GET",
}

var GetApiV4ProjectsIssuesLinksByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/links",
	Method:  "GET",
}

var PostApiV4ProjectsIssuesLinksByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/links",
	Method:  "POST",
}

var DeleteApiV4ProjectsIssuesLinksByIdByIssueIidByIssueLinkId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/links/{issue_link_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsIssuesLinksByIdByIssueIidByIssueLinkId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/links/{issue_link_id}",
	Method:  "GET",
}

var GetApiV4ProjectsIssuesMetricImagesByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/metric_images",
	Method:  "GET",
}

var PostApiV4ProjectsIssuesMetricImagesByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/metric_images",
	Method:  "POST",
}

var PostApiV4ProjectsIssuesMetricImagesAuthorizeByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/metric_images/authorize",
	Method:  "POST",
}

var DeleteApiV4ProjectsIssuesMetricImagesByIdByIssueIidByMetricImageId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/metric_images/{metric_image_id}",
	Method:  "DELETE",
}

var PutApiV4ProjectsIssuesMetricImagesByIdByIssueIidByMetricImageId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/metric_images/{metric_image_id}",
	Method:  "PUT",
}

var PostApiV4ProjectsIssuesMoveByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/move",
	Method:  "POST",
}

var GetApiV4ProjectsIssuesNotesAwardEmojiByIdByIssueIidByNoteId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/notes/{note_id}/award_emoji",
	Method:  "GET",
}

var PostApiV4ProjectsIssuesNotesAwardEmojiByIdByIssueIidByNoteId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/notes/{note_id}/award_emoji",
	Method:  "POST",
}

var DeleteApiV4ProjectsIssuesNotesAwardEmojiByIdByIssueIidByNoteIdByAwardId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/notes/{note_id}/award_emoji/{award_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsIssuesNotesAwardEmojiByIdByIssueIidByNoteIdByAwardId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/notes/{note_id}/award_emoji/{award_id}",
	Method:  "GET",
}

var GetApiV4ProjectsIssuesParticipantsByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/participants",
	Method:  "GET",
}

var GetApiV4ProjectsIssuesRelatedMergeRequestsByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/related_merge_requests",
	Method:  "GET",
}

var PutApiV4ProjectsIssuesReorderByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/reorder",
	Method:  "PUT",
}

var PostApiV4ProjectsIssuesResetSpentTimeByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/reset_spent_time",
	Method:  "POST",
}

var PostApiV4ProjectsIssuesResetTimeEstimateByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/reset_time_estimate",
	Method:  "POST",
}

var PostApiV4ProjectsIssuesTimeEstimateByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/time_estimate",
	Method:  "POST",
}

var GetApiV4ProjectsIssuesTimeStatsByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/time_stats",
	Method:  "GET",
}

var GetApiV4ProjectsIssuesUserAgentDetailByIdByIssueIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues/{issue_iid}/user_agent_detail",
	Method:  "GET",
}

var GetApiV4ProjectsIssuesStatisticsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/issues_statistics",
	Method:  "GET",
}

var GetApiV4ProjectsJobTokenScopeById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/job_token_scope",
	Method:  "GET",
}

var PatchApiV4ProjectsJobTokenScopeById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/job_token_scope",
	Method:  "PATCH",
}

var GetApiV4ProjectsJobTokenScopeAllowlistById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/job_token_scope/allowlist",
	Method:  "GET",
}

var PostApiV4ProjectsJobTokenScopeAllowlistById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/job_token_scope/allowlist",
	Method:  "POST",
}

var DeleteApiV4ProjectsJobTokenScopeAllowlistByIdByTargetProjectId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/job_token_scope/allowlist/{target_project_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsJobTokenScopeGroupsAllowlistById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/job_token_scope/groups_allowlist",
	Method:  "GET",
}

var PostApiV4ProjectsJobTokenScopeGroupsAllowlistById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/job_token_scope/groups_allowlist",
	Method:  "POST",
}

var DeleteApiV4ProjectsJobTokenScopeGroupsAllowlistByIdByTargetGroupId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/job_token_scope/groups_allowlist/{target_group_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsJobsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/jobs",
	Method:  "GET",
}

var GetApiV4ProjectsJobsArtifactsDownloadByIdByRefName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/jobs/artifacts/{ref_name}/download",
	Method:  "GET",
}

var GetApiV4ProjectsJobsArtifactsRawByIdByRefNameByArtifactPath EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/jobs/artifacts/{ref_name}/raw/{artifact_path:.+}",
	Method:  "GET",
}

var GetApiV4ProjectsJobsByIdByJobId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/jobs/{job_id}",
	Method:  "GET",
}

var DeleteApiV4ProjectsJobsArtifactsByIdByJobId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/jobs/{job_id}/artifacts",
	Method:  "DELETE",
}

var GetApiV4ProjectsJobsArtifactsByIdByJobId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/jobs/{job_id}/artifacts",
	Method:  "GET",
}

var GetApiV4ProjectsJobsArtifactsByIdByJobIdByArtifactPath EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/jobs/{job_id}/artifacts/{artifact_path:.+}",
	Method:  "GET",
}

var PostApiV4ProjectsJobsArtifactsKeepByIdByJobId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/jobs/{job_id}/artifacts/keep",
	Method:  "POST",
}

var PostApiV4ProjectsJobsCancelByIdByJobId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/jobs/{job_id}/cancel",
	Method:  "POST",
}

var PostApiV4ProjectsJobsEraseByIdByJobId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/jobs/{job_id}/erase",
	Method:  "POST",
}

var PostApiV4ProjectsJobsPlayByIdByJobId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/jobs/{job_id}/play",
	Method:  "POST",
}

var PostApiV4ProjectsJobsRetryByIdByJobId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/jobs/{job_id}/retry",
	Method:  "POST",
}

var GetApiV4ProjectsJobsTraceByIdByJobId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/jobs/{job_id}/trace",
	Method:  "GET",
}

var GetApiV4ProjectsLanguagesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/languages",
	Method:  "GET",
}

var GetApiV4ProjectsMembersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/members",
	Method:  "GET",
}

var PostApiV4ProjectsMembersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/members",
	Method:  "POST",
}

var GetApiV4ProjectsMembersAllById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/members/all",
	Method:  "GET",
}

var GetApiV4ProjectsMembersAllByIdByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/members/all/{user_id}",
	Method:  "GET",
}

var DeleteApiV4ProjectsMembersByIdByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/members/{user_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsMembersByIdByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/members/{user_id}",
	Method:  "GET",
}

var PutApiV4ProjectsMembersByIdByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/members/{user_id}",
	Method:  "PUT",
}

var GetApiV4ProjectsMergeRequestsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests",
	Method:  "GET",
}

var PostApiV4ProjectsMergeRequestsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests",
	Method:  "POST",
}

var GetApiV4ProjectsMergeRequestsResourceMilestoneEventsByIdByEventableId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{eventable_id}/resource_milestone_events",
	Method:  "GET",
}

var GetApiV4ProjectsMergeRequestsResourceMilestoneEventsByIdByEventableIdByEventId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{eventable_id}/resource_milestone_events/{event_id}",
	Method:  "GET",
}

var DeleteApiV4ProjectsMergeRequestsByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}",
	Method:  "DELETE",
}

var GetApiV4ProjectsMergeRequestsByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}",
	Method:  "GET",
}

var PutApiV4ProjectsMergeRequestsByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}",
	Method:  "PUT",
}

var PostApiV4ProjectsMergeRequestsAddSpentTimeByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/add_spent_time",
	Method:  "POST",
}

var GetApiV4ProjectsMergeRequestsApprovalStateByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/approval_state",
	Method:  "GET",
}

var GetApiV4ProjectsMergeRequestsApprovalsByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/approvals",
	Method:  "GET",
}

var PostApiV4ProjectsMergeRequestsApprovalsByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/approvals",
	Method:  "POST",
}

var PostApiV4ProjectsMergeRequestsApproveByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/approve",
	Method:  "POST",
}

var GetApiV4ProjectsMergeRequestsAwardEmojiByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/award_emoji",
	Method:  "GET",
}

var PostApiV4ProjectsMergeRequestsAwardEmojiByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/award_emoji",
	Method:  "POST",
}

var DeleteApiV4ProjectsMergeRequestsAwardEmojiByIdByMergeRequestIidByAwardId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/award_emoji/{award_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsMergeRequestsAwardEmojiByIdByMergeRequestIidByAwardId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/award_emoji/{award_id}",
	Method:  "GET",
}

var PostApiV4ProjectsMergeRequestsCancelMergeWhenPipelineSucceedsByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/cancel_merge_when_pipeline_succeeds",
	Method:  "POST",
}

var GetApiV4ProjectsMergeRequestsChangesByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/changes",
	Method:  "GET",
}

var GetApiV4ProjectsMergeRequestsClosesIssuesByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/closes_issues",
	Method:  "GET",
}

var GetApiV4ProjectsMergeRequestsCommitsByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/commits",
	Method:  "GET",
}

var DeleteApiV4ProjectsMergeRequestsContextCommitsByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/context_commits",
	Method:  "DELETE",
}

var GetApiV4ProjectsMergeRequestsContextCommitsByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/context_commits",
	Method:  "GET",
}

var PostApiV4ProjectsMergeRequestsContextCommitsByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/context_commits",
	Method:  "POST",
}

var GetApiV4ProjectsMergeRequestsDiffsByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/diffs",
	Method:  "GET",
}

var GetApiV4ProjectsMergeRequestsDraftNotesByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/draft_notes",
	Method:  "GET",
}

var PostApiV4ProjectsMergeRequestsDraftNotesByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/draft_notes",
	Method:  "POST",
}

var PostApiV4ProjectsMergeRequestsDraftNotesBulkPublishByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/draft_notes/bulk_publish",
	Method:  "POST",
}

var DeleteApiV4ProjectsMergeRequestsDraftNotesByIdByMergeRequestIidByDraftNoteId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/draft_notes/{draft_note_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsMergeRequestsDraftNotesByIdByMergeRequestIidByDraftNoteId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/draft_notes/{draft_note_id}",
	Method:  "GET",
}

var PutApiV4ProjectsMergeRequestsDraftNotesByIdByMergeRequestIidByDraftNoteId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/draft_notes/{draft_note_id}",
	Method:  "PUT",
}

var PutApiV4ProjectsMergeRequestsDraftNotesPublishByIdByMergeRequestIidByDraftNoteId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/draft_notes/{draft_note_id}/publish",
	Method:  "PUT",
}

var PutApiV4ProjectsMergeRequestsMergeByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/merge",
	Method:  "PUT",
}

var GetApiV4ProjectsMergeRequestsMergeRefByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/merge_ref",
	Method:  "GET",
}

var GetApiV4ProjectsMergeRequestsNotesAwardEmojiByIdByMergeRequestIidByNoteId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/notes/{note_id}/award_emoji",
	Method:  "GET",
}

var PostApiV4ProjectsMergeRequestsNotesAwardEmojiByIdByMergeRequestIidByNoteId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/notes/{note_id}/award_emoji",
	Method:  "POST",
}

var DeleteApiV4ProjectsMergeRequestsNotesAwardEmojiByIdByMergeRequestIidByNoteIdByAwardId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/notes/{note_id}/award_emoji/{award_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsMergeRequestsNotesAwardEmojiByIdByMergeRequestIidByNoteIdByAwardId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/notes/{note_id}/award_emoji/{award_id}",
	Method:  "GET",
}

var GetApiV4ProjectsMergeRequestsParticipantsByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/participants",
	Method:  "GET",
}

var GetApiV4ProjectsMergeRequestsPipelinesByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/pipelines",
	Method:  "GET",
}

var PostApiV4ProjectsMergeRequestsPipelinesByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/pipelines",
	Method:  "POST",
}

var GetApiV4ProjectsMergeRequestsRawDiffsByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/raw_diffs",
	Method:  "GET",
}

var PutApiV4ProjectsMergeRequestsRebaseByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/rebase",
	Method:  "PUT",
}

var GetApiV4ProjectsMergeRequestsRelatedIssuesByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/related_issues",
	Method:  "GET",
}

var PutApiV4ProjectsMergeRequestsResetApprovalsByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/reset_approvals",
	Method:  "PUT",
}

var PostApiV4ProjectsMergeRequestsResetSpentTimeByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/reset_spent_time",
	Method:  "POST",
}

var PostApiV4ProjectsMergeRequestsResetTimeEstimateByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/reset_time_estimate",
	Method:  "POST",
}

var GetApiV4ProjectsMergeRequestsReviewersByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/reviewers",
	Method:  "GET",
}

var PostApiV4ProjectsMergeRequestsTimeEstimateByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/time_estimate",
	Method:  "POST",
}

var GetApiV4ProjectsMergeRequestsTimeStatsByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/time_stats",
	Method:  "GET",
}

var PostApiV4ProjectsMergeRequestsUnapproveByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/unapprove",
	Method:  "POST",
}

var GetApiV4ProjectsMergeRequestsVersionsByIdByMergeRequestIid EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/versions",
	Method:  "GET",
}

var GetApiV4ProjectsMergeRequestsVersionsByIdByMergeRequestIidByVersionId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/merge_requests/{merge_request_iid}/versions/{version_id}",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesCargoConfigJsonById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/cargo/config.json",
	Method:  "GET",
}

var PostApiV4ProjectsPackagesComposerById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/composer",
	Method:  "POST",
}

var GetApiV4ProjectsPackagesComposerArchivesByIdByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/composer/archives/{package_name:.+}",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesConanV1ConansSearchById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/conans/search",
	Method:  "GET",
}

var DeleteApiV4ProjectsPackagesConanV1ConansByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}",
	Method:  "DELETE",
}

var GetApiV4ProjectsPackagesConanV1ConansByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesConanV1ConansDigestByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/digest",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesConanV1ConansDownloadUrlsByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/download_urls",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesConanV1ConansPackagesByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByConanPackageReference EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/packages/{conan_package_reference}",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesConanV1ConansPackagesDigestByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByConanPackageReference EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/packages/{conan_package_reference}/digest",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesConanV1ConansPackagesDownloadUrlsByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByConanPackageReference EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/packages/{conan_package_reference}/download_urls",
	Method:  "GET",
}

var PostApiV4ProjectsPackagesConanV1ConansPackagesUploadUrlsByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByConanPackageReference EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/packages/{conan_package_reference}/upload_urls",
	Method:  "POST",
}

var GetApiV4ProjectsPackagesConanV1ConansSearchByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/search",
	Method:  "GET",
}

var PostApiV4ProjectsPackagesConanV1ConansUploadUrlsByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/upload_urls",
	Method:  "POST",
}

var GetApiV4ProjectsPackagesConanV1FilesExportByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/export/{file_name}",
	Method:  "GET",
}

var PutApiV4ProjectsPackagesConanV1FilesExportByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/export/{file_name}",
	Method:  "PUT",
}

var PutApiV4ProjectsPackagesConanV1FilesExportAuthorizeByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/export/{file_name}/authorize",
	Method:  "PUT",
}

var GetApiV4ProjectsPackagesConanV1FilesPackageByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByConanPackageReferenceByPackageRevisionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/package/{conan_package_reference}/{package_revision}/{file_name}",
	Method:  "GET",
}

var PutApiV4ProjectsPackagesConanV1FilesPackageByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByConanPackageReferenceByPackageRevisionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/package/{conan_package_reference}/{package_revision}/{file_name}",
	Method:  "PUT",
}

var PutApiV4ProjectsPackagesConanV1FilesPackageAuthorizeByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByConanPackageReferenceByPackageRevisionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/package/{conan_package_reference}/{package_revision}/{file_name}/authorize",
	Method:  "PUT",
}

var GetApiV4ProjectsPackagesConanV1PingById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/ping",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesConanV1UsersAuthenticateById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/users/authenticate",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesConanV1UsersCheckCredentialsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v1/users/check_credentials",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesConanV2ConansSearchById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/conans/search",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesConanV2ConansLatestByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/conans/{package_name}/{package_version}/{package_username}/{package_channel}/latest",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesConanV2ConansRevisionsByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/conans/{package_name}/{package_version}/{package_username}/{package_channel}/revisions",
	Method:  "GET",
}

var DeleteApiV4ProjectsPackagesConanV2ConansRevisionsByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevision EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/conans/{package_name}/{package_version}/{package_username}/{package_channel}/revisions/{recipe_revision}",
	Method:  "DELETE",
}

var GetApiV4ProjectsPackagesConanV2ConansRevisionsFilesByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevision EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/conans/{package_name}/{package_version}/{package_username}/{package_channel}/revisions/{recipe_revision}/files",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesConanV2ConansRevisionsFilesByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/conans/{package_name}/{package_version}/{package_username}/{package_channel}/revisions/{recipe_revision}/files/{file_name}",
	Method:  "GET",
}

var PutApiV4ProjectsPackagesConanV2ConansRevisionsFilesByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/conans/{package_name}/{package_version}/{package_username}/{package_channel}/revisions/{recipe_revision}/files/{file_name}",
	Method:  "PUT",
}

var PutApiV4ProjectsPackagesConanV2ConansRevisionsFilesAuthorizeByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/conans/{package_name}/{package_version}/{package_username}/{package_channel}/revisions/{recipe_revision}/files/{file_name}/authorize",
	Method:  "PUT",
}

var GetApiV4ProjectsPackagesConanV2ConansRevisionsPackagesLatestByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByConanPackageReference EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/conans/{package_name}/{package_version}/{package_username}/{package_channel}/revisions/{recipe_revision}/packages/{conan_package_reference}/latest",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesConanV2ConansRevisionsPackagesRevisionsByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByConanPackageReference EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/conans/{package_name}/{package_version}/{package_username}/{package_channel}/revisions/{recipe_revision}/packages/{conan_package_reference}/revisions",
	Method:  "GET",
}

var DeleteApiV4ProjectsPackagesConanV2ConansRevisionsPackagesRevisionsByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByConanPackageReferenceByPackageRevision EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/conans/{package_name}/{package_version}/{package_username}/{package_channel}/revisions/{recipe_revision}/packages/{conan_package_reference}/revisions/{package_revision}",
	Method:  "DELETE",
}

var GetApiV4ProjectsPackagesConanV2ConansRevisionsPackagesRevisionsFilesByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByConanPackageReferenceByPackageRevision EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/conans/{package_name}/{package_version}/{package_username}/{package_channel}/revisions/{recipe_revision}/packages/{conan_package_reference}/revisions/{package_revision}/files",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesConanV2ConansRevisionsPackagesRevisionsFilesByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByConanPackageReferenceByPackageRevisionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/conans/{package_name}/{package_version}/{package_username}/{package_channel}/revisions/{recipe_revision}/packages/{conan_package_reference}/revisions/{package_revision}/files/{file_name}",
	Method:  "GET",
}

var PutApiV4ProjectsPackagesConanV2ConansRevisionsPackagesRevisionsFilesByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByConanPackageReferenceByPackageRevisionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/conans/{package_name}/{package_version}/{package_username}/{package_channel}/revisions/{recipe_revision}/packages/{conan_package_reference}/revisions/{package_revision}/files/{file_name}",
	Method:  "PUT",
}

var PutApiV4ProjectsPackagesConanV2ConansRevisionsPackagesRevisionsFilesAuthorizeByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevisionByConanPackageReferenceByPackageRevisionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/conans/{package_name}/{package_version}/{package_username}/{package_channel}/revisions/{recipe_revision}/packages/{conan_package_reference}/revisions/{package_revision}/files/{file_name}/authorize",
	Method:  "PUT",
}

var GetApiV4ProjectsPackagesConanV2ConansRevisionsSearchByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannelByRecipeRevision EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/conans/{package_name}/{package_version}/{package_username}/{package_channel}/revisions/{recipe_revision}/search",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesConanV2ConansSearchByIdByPackageNameByPackageVersionByPackageUsernameByPackageChannel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/conans/{package_name}/{package_version}/{package_username}/{package_channel}/search",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesConanV2UsersAuthenticateById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/users/authenticate",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesConanV2UsersCheckCredentialsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/conan/v2/users/check_credentials",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesDebianDistsInReleaseByIdByDistribution EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/debian/dists/{distribution:.+}/InRelease",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesDebianDistsReleaseByIdByDistribution EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/debian/dists/{distribution:.+}/Release",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesDebianDistsReleaseGpgByIdByDistribution EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/debian/dists/{distribution:.+}/Release.gpg",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesDebianDistsBinaryPackagesByIdByDistributionByComponentByArchitecture EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/debian/dists/{distribution:.+}/{component}/binary-{architecture}/Packages",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesDebianDistsBinaryByHashSHA256ByIdByDistributionByComponentByArchitectureByFileSha256 EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/debian/dists/{distribution:.+}/{component}/binary-{architecture}/by-hash/SHA256/{file_sha256}",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesDebianDistsDebianInstallerBinaryPackagesByIdByDistributionByComponentByArchitecture EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/debian/dists/{distribution:.+}/{component}/debian-installer/binary-{architecture}/Packages",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesDebianDistsDebianInstallerBinaryByHashSHA256ByIdByDistributionByComponentByArchitectureByFileSha256 EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/debian/dists/{distribution:.+}/{component}/debian-installer/binary-{architecture}/by-hash/SHA256/{file_sha256}",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesDebianDistsSourceSourcesByIdByDistributionByComponent EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/debian/dists/{distribution:.+}/{component}/source/Sources",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesDebianDistsSourceByHashSHA256ByIdByDistributionByComponentByFileSha256 EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/debian/dists/{distribution:.+}/{component}/source/by-hash/SHA256/{file_sha256}",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesDebianPoolByIdByDistributionByLetterByPackageNameByPackageVersionByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/debian/pool/{distribution}/{letter}/{package_name}/{package_version}/{file_name}",
	Method:  "GET",
}

var PutApiV4ProjectsPackagesDebianByIdByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/debian/{file_name}",
	Method:  "PUT",
}

var PutApiV4ProjectsPackagesDebianAuthorizeByIdByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/debian/{file_name}/authorize",
	Method:  "PUT",
}

var GetApiV4ProjectsPackagesGenericPackageVersionPathByIdByPackageNameByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/generic/{package_name}/*package_version/(*path/){file_name}",
	Method:  "GET",
}

var PutApiV4ProjectsPackagesGenericPackageVersionPathByIdByPackageNameByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/generic/{package_name}/*package_version/(*path/){file_name}",
	Method:  "PUT",
}

var PutApiV4ProjectsPackagesGenericPackageVersionPathAuthorizeByIdByPackageNameByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/generic/{package_name}/*package_version/(*path/){file_name}/authorize",
	Method:  "PUT",
}

var GetApiV4ProjectsPackagesGoVListByIdByModuleName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/go/{module_name:.+}/@v/list",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesGoVInfoByIdByModuleNameByModuleVersion EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/go/{module_name:.+}/@v/{module_version}.info",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesGoVModByIdByModuleNameByModuleVersion EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/go/{module_name:.+}/@v/{module_version}.mod",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesGoVZipByIdByModuleNameByModuleVersion EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/go/{module_name:.+}/@v/{module_version}.zip",
	Method:  "GET",
}

var PostApiV4ProjectsPackagesHelmApiChartsByIdByChannel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/helm/api/{channel}/charts",
	Method:  "POST",
}

var PostApiV4ProjectsPackagesHelmApiChartsAuthorizeByIdByChannel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/helm/api/{channel}/charts/authorize",
	Method:  "POST",
}

var GetApiV4ProjectsPackagesHelmChartsTgzByIdByChannelByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/helm/{channel}/charts/{file_name}.tgz",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesHelmIndexYamlByIdByChannel EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/helm/{channel}/index.yaml",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesMavenByIdByPathByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/maven/{path:.+}/{file_name}",
	Method:  "GET",
}

var PutApiV4ProjectsPackagesMavenByIdByPathByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/maven/{path:.+}/{file_name}",
	Method:  "PUT",
}

var PutApiV4ProjectsPackagesMavenAuthorizeByIdByPathByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/maven/{path:.+}/{file_name}/authorize",
	Method:  "PUT",
}

var GetApiV4ProjectsPackagesMlModelsFilesPathByIdByModelVersionIdByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/ml_models/{model_version_id}/files/(*path/){file_name}",
	Method:  "GET",
}

var PutApiV4ProjectsPackagesMlModelsFilesPathByIdByModelVersionIdByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/ml_models/{model_version_id}/files/(*path/){file_name}",
	Method:  "PUT",
}

var PutApiV4ProjectsPackagesMlModelsFilesPathAuthorizeByIdByModelVersionIdByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/ml_models/{model_version_id}/files/(*path/){file_name}/authorize",
	Method:  "PUT",
}

var GetApiV4ProjectsPackagesNpmByIdByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/npm/{package_name:.+}",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesNpmByIdByPackageNameByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/npm/{package_name:.+}/-/{file_name:.+}",
	Method:  "GET",
}

var PostApiV4ProjectsPackagesNpmNpmV1SecurityAdvisoriesBulkById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/npm/-/npm/v1/security/advisories/bulk",
	Method:  "POST",
}

var PostApiV4ProjectsPackagesNpmNpmV1SecurityAuditsQuickById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/npm/-/npm/v1/security/audits/quick",
	Method:  "POST",
}

var GetApiV4ProjectsPackagesNpmPackageDistTagsByIdByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/npm/-/package/{package_name:.+}/dist-tags",
	Method:  "GET",
}

var DeleteApiV4ProjectsPackagesNpmPackageDistTagsByIdByPackageNameByTag EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/npm/-/package/{package_name:.+}/dist-tags/{tag}",
	Method:  "DELETE",
}

var PutApiV4ProjectsPackagesNpmPackageDistTagsByIdByPackageNameByTag EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/npm/-/package/{package_name:.+}/dist-tags/{tag}",
	Method:  "PUT",
}

var PutApiV4ProjectsPackagesNpmByIdByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/npm/{package_name}",
	Method:  "PUT",
}

var PutApiV4ProjectsPackagesNugetById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/nuget",
	Method:  "PUT",
}

var DeleteApiV4ProjectsPackagesNugetByIdByPackageNameByPackageVersion EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/nuget/{package_name:.+}/{package_version:.+}",
	Method:  "DELETE",
}

var PutApiV4ProjectsPackagesNugetAuthorizeById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/nuget/authorize",
	Method:  "PUT",
}

var GetApiV4ProjectsPackagesNugetDownloadByIdByPackageNameByPackageVersionByPackageFilename EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/nuget/download/{package_name:.+}/{package_version:.+}/{package_filename:.+}",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesNugetDownloadIndexByIdByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/nuget/download/{package_name:.+}/index",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesNugetIndexById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/nuget/index",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesNugetMetadataByIdByPackageNameByPackageVersion EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/nuget/metadata/{package_name:.+}/{package_version:.+}",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesNugetMetadataIndexByIdByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/nuget/metadata/{package_name:.+}/index",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesNugetQueryById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/nuget/query",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesNugetSymbolfilesFileNameSignatureSameFileNameById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/nuget/symbolfiles/*file_name/*signature/*same_file_name",
	Method:  "GET",
}

var PutApiV4ProjectsPackagesNugetSymbolpackageById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/nuget/symbolpackage",
	Method:  "PUT",
}

var PutApiV4ProjectsPackagesNugetSymbolpackageAuthorizeById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/nuget/symbolpackage/authorize",
	Method:  "PUT",
}

var GetApiV4ProjectsPackagesNugetV2ById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/nuget/v2",
	Method:  "GET",
}

var PutApiV4ProjectsPackagesNugetV2ById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/nuget/v2",
	Method:  "PUT",
}

var GetApiV4ProjectsPackagesNugetV2MetadataById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/nuget/v2/$metadata",
	Method:  "GET",
}

var PutApiV4ProjectsPackagesNugetV2AuthorizeById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/nuget/v2/authorize",
	Method:  "PUT",
}

var GetApiV4ProjectsPackagesProtectionRulesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/protection/rules",
	Method:  "GET",
}

var PostApiV4ProjectsPackagesProtectionRulesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/protection/rules",
	Method:  "POST",
}

var DeleteApiV4ProjectsPackagesProtectionRulesByIdByPackageProtectionRuleId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/protection/rules/{package_protection_rule_id}",
	Method:  "DELETE",
}

var PatchApiV4ProjectsPackagesProtectionRulesByIdByPackageProtectionRuleId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/protection/rules/{package_protection_rule_id}",
	Method:  "PATCH",
}

var PostApiV4ProjectsPackagesPypiById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/pypi",
	Method:  "POST",
}

var PostApiV4ProjectsPackagesPypiAuthorizeById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/pypi/authorize",
	Method:  "POST",
}

var GetApiV4ProjectsPackagesPypiFilesByIdBySha256ByFileIdentifier EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/pypi/files/{sha256}/{file_identifier:.+}",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesPypiSimpleById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/pypi/simple",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesPypiSimpleByIdByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/pypi/simple/{package_name:.+}",
	Method:  "GET",
}

var PostApiV4ProjectsPackagesRpmById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/rpm",
	Method:  "POST",
}

var GetApiV4ProjectsPackagesRpmByIdByPackageFileIdByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/rpm/{package_file_id:.+}/{file_name:.+}",
	Method:  "GET",
}

var PostApiV4ProjectsPackagesRpmAuthorizeById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/rpm/authorize",
	Method:  "POST",
}

var GetApiV4ProjectsPackagesRpmRepodataByIdByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/rpm/repodata/{file_name:.+}",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesRubygemsApiV1DependenciesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/rubygems/api/v1/dependencies",
	Method:  "GET",
}

var PostApiV4ProjectsPackagesRubygemsApiV1GemsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/rubygems/api/v1/gems",
	Method:  "POST",
}

var PostApiV4ProjectsPackagesRubygemsApiV1GemsAuthorizeById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/rubygems/api/v1/gems/authorize",
	Method:  "POST",
}

var GetApiV4ProjectsPackagesRubygemsGemsByIdByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/rubygems/gems/{file_name}",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesRubygemsQuickMarshal48ByIdByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/rubygems/quick/Marshal.4.8/{file_name}",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesRubygemsByIdByFileName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/rubygems/{file_name}",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesTerraformModulesByIdByModuleNameByModuleSystem EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/terraform/modules/{module_name}/{module_system}",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesTerraformModulesModuleVersionByIdByModuleNameByModuleSystem EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/terraform/modules/{module_name}/{module_system}/*module_version",
	Method:  "GET",
}

var PutApiV4ProjectsPackagesTerraformModulesFileByIdByModuleNameByModuleSystemByModuleVersion EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/terraform/modules/{module_name}/{module_system}/{module_version:.+}/file",
	Method:  "PUT",
}

var PutApiV4ProjectsPackagesTerraformModulesFileAuthorizeByIdByModuleNameByModuleSystemByModuleVersion EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/terraform/modules/{module_name}/{module_system}/{module_version:.+}/file/authorize",
	Method:  "PUT",
}

var DeleteApiV4ProjectsPackagesByIdByPackageId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/{package_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsPackagesByIdByPackageId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/{package_id}",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesPackageFilesByIdByPackageId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/{package_id}/package_files",
	Method:  "GET",
}

var DeleteApiV4ProjectsPackagesPackageFilesByIdByPackageIdByPackageFileId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/{package_id}/package_files/{package_file_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsPackagesPipelinesByIdByPackageId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/packages/{package_id}/pipelines",
	Method:  "GET",
}

var DeleteApiV4ProjectsPagesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pages",
	Method:  "DELETE",
}

var GetApiV4ProjectsPagesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pages",
	Method:  "GET",
}

var PatchApiV4ProjectsPagesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pages",
	Method:  "PATCH",
}

var GetApiV4ProjectsPagesDomainsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pages/domains",
	Method:  "GET",
}

var PostApiV4ProjectsPagesDomainsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pages/domains",
	Method:  "POST",
}

var DeleteApiV4ProjectsPagesDomainsByIdByDomain EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pages/domains/{domain}",
	Method:  "DELETE",
}

var GetApiV4ProjectsPagesDomainsByIdByDomain EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pages/domains/{domain}",
	Method:  "GET",
}

var PutApiV4ProjectsPagesDomainsByIdByDomain EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pages/domains/{domain}",
	Method:  "PUT",
}

var PutApiV4ProjectsPagesDomainsVerifyByIdByDomain EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pages/domains/{domain}/verify",
	Method:  "PUT",
}

var GetApiV4ProjectsPagesAccessById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pages_access",
	Method:  "GET",
}

var PostApiV4ProjectsPipelineById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipeline",
	Method:  "POST",
}

var GetApiV4ProjectsPipelineSchedulesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipeline_schedules",
	Method:  "GET",
}

var PostApiV4ProjectsPipelineSchedulesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipeline_schedules",
	Method:  "POST",
}

var DeleteApiV4ProjectsPipelineSchedulesByIdByPipelineScheduleId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipeline_schedules/{pipeline_schedule_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsPipelineSchedulesByIdByPipelineScheduleId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipeline_schedules/{pipeline_schedule_id}",
	Method:  "GET",
}

var PutApiV4ProjectsPipelineSchedulesByIdByPipelineScheduleId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipeline_schedules/{pipeline_schedule_id}",
	Method:  "PUT",
}

var GetApiV4ProjectsPipelineSchedulesPipelinesByIdByPipelineScheduleId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipeline_schedules/{pipeline_schedule_id}/pipelines",
	Method:  "GET",
}

var PostApiV4ProjectsPipelineSchedulesPlayByIdByPipelineScheduleId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipeline_schedules/{pipeline_schedule_id}/play",
	Method:  "POST",
}

var PostApiV4ProjectsPipelineSchedulesTakeOwnershipByIdByPipelineScheduleId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipeline_schedules/{pipeline_schedule_id}/take_ownership",
	Method:  "POST",
}

var PostApiV4ProjectsPipelineSchedulesVariablesByIdByPipelineScheduleId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipeline_schedules/{pipeline_schedule_id}/variables",
	Method:  "POST",
}

var DeleteApiV4ProjectsPipelineSchedulesVariablesByIdByPipelineScheduleIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipeline_schedules/{pipeline_schedule_id}/variables/{key}",
	Method:  "DELETE",
}

var PutApiV4ProjectsPipelineSchedulesVariablesByIdByPipelineScheduleIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipeline_schedules/{pipeline_schedule_id}/variables/{key}",
	Method:  "PUT",
}

var GetApiV4ProjectsPipelinesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipelines",
	Method:  "GET",
}

var GetApiV4ProjectsPipelinesLatestById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipelines/latest",
	Method:  "GET",
}

var DeleteApiV4ProjectsPipelinesByIdByPipelineId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipelines/{pipeline_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsPipelinesByIdByPipelineId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipelines/{pipeline_id}",
	Method:  "GET",
}

var GetApiV4ProjectsPipelinesBridgesByIdByPipelineId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipelines/{pipeline_id}/bridges",
	Method:  "GET",
}

var PostApiV4ProjectsPipelinesCancelByIdByPipelineId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipelines/{pipeline_id}/cancel",
	Method:  "POST",
}

var GetApiV4ProjectsPipelinesJobsByIdByPipelineId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipelines/{pipeline_id}/jobs",
	Method:  "GET",
}

var PutApiV4ProjectsPipelinesMetadataByIdByPipelineId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipelines/{pipeline_id}/metadata",
	Method:  "PUT",
}

var PostApiV4ProjectsPipelinesRetryByIdByPipelineId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipelines/{pipeline_id}/retry",
	Method:  "POST",
}

var GetApiV4ProjectsPipelinesTestReportByIdByPipelineId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipelines/{pipeline_id}/test_report",
	Method:  "GET",
}

var GetApiV4ProjectsPipelinesTestReportSummaryByIdByPipelineId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipelines/{pipeline_id}/test_report_summary",
	Method:  "GET",
}

var GetApiV4ProjectsPipelinesVariablesByIdByPipelineId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/pipelines/{pipeline_id}/variables",
	Method:  "GET",
}

var GetApiV4ProjectsProtectedBranchesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/protected_branches",
	Method:  "GET",
}

var PostApiV4ProjectsProtectedBranchesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/protected_branches",
	Method:  "POST",
}

var DeleteApiV4ProjectsProtectedBranchesByIdByName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/protected_branches/{name}",
	Method:  "DELETE",
}

var GetApiV4ProjectsProtectedBranchesByIdByName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/protected_branches/{name}",
	Method:  "GET",
}

var PatchApiV4ProjectsProtectedBranchesByIdByName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/protected_branches/{name}",
	Method:  "PATCH",
}

var GetApiV4ProjectsProtectedTagsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/protected_tags",
	Method:  "GET",
}

var PostApiV4ProjectsProtectedTagsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/protected_tags",
	Method:  "POST",
}

var DeleteApiV4ProjectsProtectedTagsByIdByName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/protected_tags/{name}",
	Method:  "DELETE",
}

var GetApiV4ProjectsProtectedTagsByIdByName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/protected_tags/{name}",
	Method:  "GET",
}

var GetApiV4ProjectsRegistryProtectionRepositoryRulesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/registry/protection/repository/rules",
	Method:  "GET",
}

var PostApiV4ProjectsRegistryProtectionRepositoryRulesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/registry/protection/repository/rules",
	Method:  "POST",
}

var DeleteApiV4ProjectsRegistryProtectionRepositoryRulesByIdByProtectionRuleId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/registry/protection/repository/rules/{protection_rule_id}",
	Method:  "DELETE",
}

var PatchApiV4ProjectsRegistryProtectionRepositoryRulesByIdByProtectionRuleId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/registry/protection/repository/rules/{protection_rule_id}",
	Method:  "PATCH",
}

var GetApiV4ProjectsRegistryRepositoriesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/registry/repositories",
	Method:  "GET",
}

var DeleteApiV4ProjectsRegistryRepositoriesByIdByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/registry/repositories/{repository_id}",
	Method:  "DELETE",
}

var DeleteApiV4ProjectsRegistryRepositoriesTagsByIdByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/registry/repositories/{repository_id}/tags",
	Method:  "DELETE",
}

var GetApiV4ProjectsRegistryRepositoriesTagsByIdByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/registry/repositories/{repository_id}/tags",
	Method:  "GET",
}

var DeleteApiV4ProjectsRegistryRepositoriesTagsByIdByRepositoryIdByTagName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/registry/repositories/{repository_id}/tags/{tag_name}",
	Method:  "DELETE",
}

var GetApiV4ProjectsRegistryRepositoriesTagsByIdByRepositoryIdByTagName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/registry/repositories/{repository_id}/tags/{tag_name}",
	Method:  "GET",
}

var GetApiV4ProjectsRelationImportsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/relation-imports",
	Method:  "GET",
}

var GetApiV4ProjectsReleasesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/releases",
	Method:  "GET",
}

var PostApiV4ProjectsReleasesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/releases",
	Method:  "POST",
}

var GetApiV4ProjectsReleasesPermalinkLatestSuffixPathById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/releases/permalink/latest(/)(*suffix_path)",
	Method:  "GET",
}

var DeleteApiV4ProjectsReleasesByIdByTagName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/releases/{tag_name}",
	Method:  "DELETE",
}

var GetApiV4ProjectsReleasesByIdByTagName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/releases/{tag_name}",
	Method:  "GET",
}

var PutApiV4ProjectsReleasesByIdByTagName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/releases/{tag_name}",
	Method:  "PUT",
}

var GetApiV4ProjectsReleasesAssetsLinksByIdByTagName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/releases/{tag_name}/assets/links",
	Method:  "GET",
}

var PostApiV4ProjectsReleasesAssetsLinksByIdByTagName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/releases/{tag_name}/assets/links",
	Method:  "POST",
}

var DeleteApiV4ProjectsReleasesAssetsLinksByIdByTagNameByLinkId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/releases/{tag_name}/assets/links/{link_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsReleasesAssetsLinksByIdByTagNameByLinkId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/releases/{tag_name}/assets/links/{link_id}",
	Method:  "GET",
}

var PutApiV4ProjectsReleasesAssetsLinksByIdByTagNameByLinkId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/releases/{tag_name}/assets/links/{link_id}",
	Method:  "PUT",
}

var GetApiV4ProjectsReleasesDownloadsByIdByTagNameByDirectAssetPath EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/releases/{tag_name}/downloads/{direct_asset_path:.+}",
	Method:  "GET",
}

var PostApiV4ProjectsReleasesEvidenceByIdByTagName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/releases/{tag_name}/evidence",
	Method:  "POST",
}

var GetApiV4ProjectsRemoteMirrorsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/remote_mirrors",
	Method:  "GET",
}

var PostApiV4ProjectsRemoteMirrorsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/remote_mirrors",
	Method:  "POST",
}

var DeleteApiV4ProjectsRemoteMirrorsByIdByMirrorId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/remote_mirrors/{mirror_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsRemoteMirrorsByIdByMirrorId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/remote_mirrors/{mirror_id}",
	Method:  "GET",
}

var PutApiV4ProjectsRemoteMirrorsByIdByMirrorId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/remote_mirrors/{mirror_id}",
	Method:  "PUT",
}

var GetApiV4ProjectsRemoteMirrorsPublicKeyByIdByMirrorId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/remote_mirrors/{mirror_id}/public_key",
	Method:  "GET",
}

var PostApiV4ProjectsRemoteMirrorsSyncByIdByMirrorId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/remote_mirrors/{mirror_id}/sync",
	Method:  "POST",
}

var GetApiV4ProjectsRepositoryArchiveById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/archive",
	Method:  "GET",
}

var GetApiV4ProjectsRepositoryBlobsByIdBySha EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/blobs/{sha}",
	Method:  "GET",
}

var GetApiV4ProjectsRepositoryBlobsRawByIdBySha EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/blobs/{sha}/raw",
	Method:  "GET",
}

var GetApiV4ProjectsRepositoryBranchesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/branches",
	Method:  "GET",
}

var PostApiV4ProjectsRepositoryBranchesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/branches",
	Method:  "POST",
}

var DeleteApiV4ProjectsRepositoryBranchesByIdByBranch EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/branches/{branch}",
	Method:  "DELETE",
}

var GetApiV4ProjectsRepositoryBranchesByIdByBranch EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/branches/{branch}",
	Method:  "GET",
}

var HeadApiV4ProjectsRepositoryBranchesByIdByBranch EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/branches/{branch}",
	Method:  "HEAD",
}

var PutApiV4ProjectsRepositoryBranchesProtectByIdByBranch EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/branches/{branch}/protect",
	Method:  "PUT",
}

var PutApiV4ProjectsRepositoryBranchesUnprotectByIdByBranch EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/branches/{branch}/unprotect",
	Method:  "PUT",
}

var GetApiV4ProjectsRepositoryChangelogById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/changelog",
	Method:  "GET",
}

var PostApiV4ProjectsRepositoryChangelogById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/changelog",
	Method:  "POST",
}

var GetApiV4ProjectsRepositoryCommitsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/commits",
	Method:  "GET",
}

var PostApiV4ProjectsRepositoryCommitsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/commits",
	Method:  "POST",
}

var GetApiV4ProjectsRepositoryCommitsByIdBySha EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/commits/{sha}",
	Method:  "GET",
}

var PostApiV4ProjectsRepositoryCommitsCherryPickByIdBySha EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/commits/{sha}/cherry_pick",
	Method:  "POST",
}

var GetApiV4ProjectsRepositoryCommitsCommentsByIdBySha EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/commits/{sha}/comments",
	Method:  "GET",
}

var PostApiV4ProjectsRepositoryCommitsCommentsByIdBySha EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/commits/{sha}/comments",
	Method:  "POST",
}

var GetApiV4ProjectsRepositoryCommitsDiffByIdBySha EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/commits/{sha}/diff",
	Method:  "GET",
}

var GetApiV4ProjectsRepositoryCommitsMergeRequestsByIdBySha EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/commits/{sha}/merge_requests",
	Method:  "GET",
}

var GetApiV4ProjectsRepositoryCommitsRefsByIdBySha EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/commits/{sha}/refs",
	Method:  "GET",
}

var PostApiV4ProjectsRepositoryCommitsRevertByIdBySha EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/commits/{sha}/revert",
	Method:  "POST",
}

var GetApiV4ProjectsRepositoryCommitsSequenceByIdBySha EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/commits/{sha}/sequence",
	Method:  "GET",
}

var GetApiV4ProjectsRepositoryCommitsSignatureByIdBySha EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/commits/{sha}/signature",
	Method:  "GET",
}

var GetApiV4ProjectsRepositoryCommitsStatusesByIdBySha EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/commits/{sha}/statuses",
	Method:  "GET",
}

var GetApiV4ProjectsRepositoryCompareById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/compare",
	Method:  "GET",
}

var GetApiV4ProjectsRepositoryContributorsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/contributors",
	Method:  "GET",
}

var DeleteApiV4ProjectsRepositoryFilesByIdByFilePath EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/files/{file_path}",
	Method:  "DELETE",
}

var GetApiV4ProjectsRepositoryFilesByIdByFilePath EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/files/{file_path}",
	Method:  "GET",
}

var HeadApiV4ProjectsRepositoryFilesByIdByFilePath EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/files/{file_path}",
	Method:  "HEAD",
}

var PostApiV4ProjectsRepositoryFilesByIdByFilePath EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/files/{file_path}",
	Method:  "POST",
}

var PutApiV4ProjectsRepositoryFilesByIdByFilePath EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/files/{file_path}",
	Method:  "PUT",
}

var GetApiV4ProjectsRepositoryFilesBlameByIdByFilePath EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/files/{file_path}/blame",
	Method:  "GET",
}

var HeadApiV4ProjectsRepositoryFilesBlameByIdByFilePath EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/files/{file_path}/blame",
	Method:  "HEAD",
}

var GetApiV4ProjectsRepositoryFilesRawByIdByFilePath EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/files/{file_path}/raw",
	Method:  "GET",
}

var GetApiV4ProjectsRepositoryHealthById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/health",
	Method:  "GET",
}

var GetApiV4ProjectsRepositoryMergeBaseById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/merge_base",
	Method:  "GET",
}

var DeleteApiV4ProjectsRepositoryMergedBranchesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/merged_branches",
	Method:  "DELETE",
}

var PutApiV4ProjectsRepositorySubmodulesByIdBySubmodule EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/submodules/{submodule}",
	Method:  "PUT",
}

var GetApiV4ProjectsRepositoryTagsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/tags",
	Method:  "GET",
}

var PostApiV4ProjectsRepositoryTagsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/tags",
	Method:  "POST",
}

var DeleteApiV4ProjectsRepositoryTagsByIdByTagName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/tags/{tag_name}",
	Method:  "DELETE",
}

var GetApiV4ProjectsRepositoryTagsByIdByTagName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/tags/{tag_name}",
	Method:  "GET",
}

var GetApiV4ProjectsRepositoryTagsSignatureByIdByTagName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/tags/{tag_name}/signature",
	Method:  "GET",
}

var GetApiV4ProjectsRepositoryTreeById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository/tree",
	Method:  "GET",
}

var PostApiV4ProjectsRepositorySizeById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/repository_size",
	Method:  "POST",
}

var GetApiV4ProjectsResourceGroupsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/resource_groups",
	Method:  "GET",
}

var GetApiV4ProjectsResourceGroupsByIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/resource_groups/{key}",
	Method:  "GET",
}

var PutApiV4ProjectsResourceGroupsByIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/resource_groups/{key}",
	Method:  "PUT",
}

var GetApiV4ProjectsResourceGroupsUpcomingJobsByIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/resource_groups/{key}/upcoming_jobs",
	Method:  "GET",
}

var PostApiV4ProjectsRestoreById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/restore",
	Method:  "POST",
}

var GetApiV4ProjectsRunnersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/runners",
	Method:  "GET",
}

var PostApiV4ProjectsRunnersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/runners",
	Method:  "POST",
}

var PostApiV4ProjectsRunnersResetRegistrationTokenById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/runners/reset_registration_token",
	Method:  "POST",
}

var DeleteApiV4ProjectsRunnersByIdByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/runners/{runner_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsSecureFilesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/secure_files",
	Method:  "GET",
}

var PostApiV4ProjectsSecureFilesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/secure_files",
	Method:  "POST",
}

var DeleteApiV4ProjectsSecureFilesByIdBySecureFileId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/secure_files/{secure_file_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsSecureFilesByIdBySecureFileId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/secure_files/{secure_file_id}",
	Method:  "GET",
}

var GetApiV4ProjectsSecureFilesDownloadByIdBySecureFileId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/secure_files/{secure_file_id}/download",
	Method:  "GET",
}

var GetApiV4ProjectsServicesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services",
	Method:  "GET",
}

var PutApiV4ProjectsServicesAppleAppStoreById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/apple-app-store",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesAsanaById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/asana",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesAssemblaById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/assembla",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesBambooById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/bamboo",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesBugzillaById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/bugzilla",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesBuildkiteById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/buildkite",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesCampfireById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/campfire",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesClickupById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/clickup",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesConfluenceById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/confluence",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesCustomIssueTrackerById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/custom-issue-tracker",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesDatadogById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/datadog",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesDiffblueCoverById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/diffblue-cover",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesDiscordById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/discord",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesDroneCiById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/drone-ci",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesEmailsOnPushById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/emails-on-push",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesEwmById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/ewm",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesExternalWikiById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/external-wiki",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesGitGuardianById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/git-guardian",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesGithubById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/github",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesGitlabSlackApplicationById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/gitlab-slack-application",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesGoogleCloudPlatformArtifactRegistryById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/google-cloud-platform-artifact-registry",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesGoogleCloudPlatformWorkloadIdentityFederationById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/google-cloud-platform-workload-identity-federation",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesGooglePlayById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/google-play",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesHangoutsChatById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/hangouts-chat",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesHarborById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/harbor",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesIrkerById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/irker",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesJenkinsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/jenkins",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesJiraById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/jira",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesJiraCloudAppById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/jira-cloud-app",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesLinearById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/linear",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesMatrixById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/matrix",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesMattermostById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/mattermost",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesMattermostSlashCommandsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/mattermost-slash-commands",
	Method:  "PUT",
}

var PostApiV4ProjectsServicesMattermostSlashCommandsTriggerById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/mattermost_slash_commands/trigger",
	Method:  "POST",
}

var PutApiV4ProjectsServicesMicrosoftTeamsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/microsoft-teams",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesMockCiById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/mock-ci",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesMockMonitoringById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/mock-monitoring",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesPackagistById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/packagist",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesPhorgeById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/phorge",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesPipelinesEmailById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/pipelines-email",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesPivotaltrackerById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/pivotaltracker",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesPrometheusById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/prometheus",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesPumbleById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/pumble",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesPushoverById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/pushover",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesRedmineById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/redmine",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesSlackById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/slack",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesSlackSlashCommandsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/slack-slash-commands",
	Method:  "PUT",
}

var PostApiV4ProjectsServicesSlackSlashCommandsTriggerById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/slack_slash_commands/trigger",
	Method:  "POST",
}

var PutApiV4ProjectsServicesSquashTmById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/squash-tm",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesTeamcityById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/teamcity",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesTelegramById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/telegram",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesUnifyCircuitById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/unify-circuit",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesWebexTeamsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/webex-teams",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesYoutrackById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/youtrack",
	Method:  "PUT",
}

var PutApiV4ProjectsServicesZentaoById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/zentao",
	Method:  "PUT",
}

var DeleteApiV4ProjectsServicesByIdBySlug EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/{slug}",
	Method:  "DELETE",
}

var GetApiV4ProjectsServicesByIdBySlug EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/services/{slug}",
	Method:  "GET",
}

var PostApiV4ProjectsShareById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/share",
	Method:  "POST",
}

var DeleteApiV4ProjectsShareByIdByGroupId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/share/{group_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsShareLocationsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/share_locations",
	Method:  "GET",
}

var GetApiV4ProjectsSnapshotById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/snapshot",
	Method:  "GET",
}

var GetApiV4ProjectsSnippetsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/snippets",
	Method:  "GET",
}

var PostApiV4ProjectsSnippetsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/snippets",
	Method:  "POST",
}

var DeleteApiV4ProjectsSnippetsByIdBySnippetId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/snippets/{snippet_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsSnippetsByIdBySnippetId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/snippets/{snippet_id}",
	Method:  "GET",
}

var PutApiV4ProjectsSnippetsByIdBySnippetId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/snippets/{snippet_id}",
	Method:  "PUT",
}

var GetApiV4ProjectsSnippetsAwardEmojiByIdBySnippetId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/snippets/{snippet_id}/award_emoji",
	Method:  "GET",
}

var PostApiV4ProjectsSnippetsAwardEmojiByIdBySnippetId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/snippets/{snippet_id}/award_emoji",
	Method:  "POST",
}

var DeleteApiV4ProjectsSnippetsAwardEmojiByIdBySnippetIdByAwardId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/snippets/{snippet_id}/award_emoji/{award_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsSnippetsAwardEmojiByIdBySnippetIdByAwardId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/snippets/{snippet_id}/award_emoji/{award_id}",
	Method:  "GET",
}

var GetApiV4ProjectsSnippetsFilesRawByIdBySnippetIdByRefByFilePath EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/snippets/{snippet_id}/files/{ref}/{file_path}/raw",
	Method:  "GET",
}

var GetApiV4ProjectsSnippetsNotesAwardEmojiByIdBySnippetIdByNoteId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/snippets/{snippet_id}/notes/{note_id}/award_emoji",
	Method:  "GET",
}

var PostApiV4ProjectsSnippetsNotesAwardEmojiByIdBySnippetIdByNoteId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/snippets/{snippet_id}/notes/{note_id}/award_emoji",
	Method:  "POST",
}

var DeleteApiV4ProjectsSnippetsNotesAwardEmojiByIdBySnippetIdByNoteIdByAwardId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/snippets/{snippet_id}/notes/{note_id}/award_emoji/{award_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsSnippetsNotesAwardEmojiByIdBySnippetIdByNoteIdByAwardId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/snippets/{snippet_id}/notes/{note_id}/award_emoji/{award_id}",
	Method:  "GET",
}

var GetApiV4ProjectsSnippetsRawByIdBySnippetId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/snippets/{snippet_id}/raw",
	Method:  "GET",
}

var GetApiV4ProjectsSnippetsUserAgentDetailByIdBySnippetId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/snippets/{snippet_id}/user_agent_detail",
	Method:  "GET",
}

var PostApiV4ProjectsStarById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/star",
	Method:  "POST",
}

var GetApiV4ProjectsStarrersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/starrers",
	Method:  "GET",
}

var GetApiV4ProjectsStatisticsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/statistics",
	Method:  "GET",
}

var PostApiV4ProjectsStatusesByIdBySha EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/statuses/{sha}",
	Method:  "POST",
}

var GetApiV4ProjectsStorageById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/storage",
	Method:  "GET",
}

var GetApiV4ProjectsTemplatesByIdByType EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/templates/{type}",
	Method:  "GET",
}

var GetApiV4ProjectsTemplatesByIdByTypeByName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/templates/{type}/{name}",
	Method:  "GET",
}

var DeleteApiV4ProjectsTerraformStateByIdByName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/terraform/state/{name}",
	Method:  "DELETE",
}

var GetApiV4ProjectsTerraformStateByIdByName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/terraform/state/{name}",
	Method:  "GET",
}

var PostApiV4ProjectsTerraformStateByIdByName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/terraform/state/{name}",
	Method:  "POST",
}

var DeleteApiV4ProjectsTerraformStateLockByIdByName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/terraform/state/{name}/lock",
	Method:  "DELETE",
}

var PostApiV4ProjectsTerraformStateLockByIdByName EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/terraform/state/{name}/lock",
	Method:  "POST",
}

var DeleteApiV4ProjectsTerraformStateVersionsByIdByNameBySerial EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/terraform/state/{name}/versions/{serial}",
	Method:  "DELETE",
}

var GetApiV4ProjectsTerraformStateVersionsByIdByNameBySerial EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/terraform/state/{name}/versions/{serial}",
	Method:  "GET",
}

var PutApiV4ProjectsTransferById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/transfer",
	Method:  "PUT",
}

var GetApiV4ProjectsTransferLocationsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/transfer_locations",
	Method:  "GET",
}

var GetApiV4ProjectsTriggersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/triggers",
	Method:  "GET",
}

var PostApiV4ProjectsTriggersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/triggers",
	Method:  "POST",
}

var DeleteApiV4ProjectsTriggersByIdByTriggerId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/triggers/{trigger_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsTriggersByIdByTriggerId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/triggers/{trigger_id}",
	Method:  "GET",
}

var PutApiV4ProjectsTriggersByIdByTriggerId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/triggers/{trigger_id}",
	Method:  "PUT",
}

var PostApiV4ProjectsUnarchiveById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/unarchive",
	Method:  "POST",
}

var PostApiV4ProjectsUnstarById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/unstar",
	Method:  "POST",
}

var GetApiV4ProjectsUploadsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/uploads",
	Method:  "GET",
}

var PostApiV4ProjectsUploadsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/uploads",
	Method:  "POST",
}

var PostApiV4ProjectsUploadsAuthorizeById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/uploads/authorize",
	Method:  "POST",
}

var DeleteApiV4ProjectsUploadsByIdBySecretByFilename EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/uploads/{secret}/{filename}",
	Method:  "DELETE",
}

var GetApiV4ProjectsUploadsByIdBySecretByFilename EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/uploads/{secret}/{filename}",
	Method:  "GET",
}

var DeleteApiV4ProjectsUploadsByIdByUploadId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/uploads/{upload_id}",
	Method:  "DELETE",
}

var GetApiV4ProjectsUploadsByIdByUploadId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/uploads/{upload_id}",
	Method:  "GET",
}

var GetApiV4ProjectsUsersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/users",
	Method:  "GET",
}

var GetApiV4ProjectsVariablesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/variables",
	Method:  "GET",
}

var PostApiV4ProjectsVariablesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/variables",
	Method:  "POST",
}

var DeleteApiV4ProjectsVariablesByIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/variables/{key}",
	Method:  "DELETE",
}

var GetApiV4ProjectsVariablesByIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/variables/{key}",
	Method:  "GET",
}

var PutApiV4ProjectsVariablesByIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/variables/{key}",
	Method:  "PUT",
}

var GetApiV4ProjectsWikisById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/wikis",
	Method:  "GET",
}

var PostApiV4ProjectsWikisById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/wikis",
	Method:  "POST",
}

var PostApiV4ProjectsWikisAttachmentsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/wikis/attachments",
	Method:  "POST",
}

var DeleteApiV4ProjectsWikisByIdBySlug EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/wikis/{slug}",
	Method:  "DELETE",
}

var GetApiV4ProjectsWikisByIdBySlug EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/wikis/{slug}",
	Method:  "GET",
}

var PutApiV4ProjectsWikisByIdBySlug EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{id}/wikis/{slug}",
	Method:  "PUT",
}

var GetApiV4ProjectsPackagesNugetV2FindPackagesByIdByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{project_id}/packages/nuget/v2/FindPackagesById()",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesNugetV2PackagesIdVersionByProjectIdByPackageNameByPackageVersion EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{project_id}/packages/nuget/v2/Packages(Id='{package_name}',Version='{package_version}')",
	Method:  "GET",
}

var GetApiV4ProjectsPackagesNugetV2PackagesByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/projects/{project_id}/packages/nuget/v2/Packages()",
	Method:  "GET",
}

var GetApiV4RegistryRepositoriesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/registry/repositories/{id}",
	Method:  "GET",
}

var DeleteApiV4Runners EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/runners",
	Method:  "DELETE",
}

var GetApiV4Runners EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/runners",
	Method:  "GET",
}

var PostApiV4Runners EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/runners",
	Method:  "POST",
}

var GetApiV4RunnersAll EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/runners/all",
	Method:  "GET",
}

var DeleteApiV4RunnersManagers EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/runners/managers",
	Method:  "DELETE",
}

var PostApiV4RunnersResetAuthenticationToken EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/runners/reset_authentication_token",
	Method:  "POST",
}

var PostApiV4RunnersResetRegistrationToken EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/runners/reset_registration_token",
	Method:  "POST",
}

var PostApiV4RunnersVerify EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/runners/verify",
	Method:  "POST",
}

var DeleteApiV4RunnersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/runners/{id}",
	Method:  "DELETE",
}

var GetApiV4RunnersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/runners/{id}",
	Method:  "GET",
}

var PutApiV4RunnersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/runners/{id}",
	Method:  "PUT",
}

var GetApiV4RunnersJobsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/runners/{id}/jobs",
	Method:  "GET",
}

var GetApiV4RunnersManagersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/runners/{id}/managers",
	Method:  "GET",
}

var PostApiV4RunnersResetAuthenticationTokenById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/runners/{id}/reset_authentication_token",
	Method:  "POST",
}

var PostApiV4SlackTrigger EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/slack/trigger",
	Method:  "POST",
}

var GetApiV4SnippetRepositoryStorageMoves EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/snippet_repository_storage_moves",
	Method:  "GET",
}

var PostApiV4SnippetRepositoryStorageMoves EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/snippet_repository_storage_moves",
	Method:  "POST",
}

var GetApiV4SnippetRepositoryStorageMovesByRepositoryStorageMoveId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/snippet_repository_storage_moves/{repository_storage_move_id}",
	Method:  "GET",
}

var GetApiV4Snippets EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/snippets",
	Method:  "GET",
}

var PostApiV4Snippets EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/snippets",
	Method:  "POST",
}

var GetApiV4SnippetsAll EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/snippets/all",
	Method:  "GET",
}

var GetApiV4SnippetsPublic EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/snippets/public",
	Method:  "GET",
}

var DeleteApiV4SnippetsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/snippets/{id}",
	Method:  "DELETE",
}

var GetApiV4SnippetsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/snippets/{id}",
	Method:  "GET",
}

var PutApiV4SnippetsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/snippets/{id}",
	Method:  "PUT",
}

var GetApiV4SnippetsFilesRawByIdByRefByFilePath EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/snippets/{id}/files/{ref}/{file_path}/raw",
	Method:  "GET",
}

var GetApiV4SnippetsRawById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/snippets/{id}/raw",
	Method:  "GET",
}

var GetApiV4SnippetsRepositoryStorageMovesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/snippets/{id}/repository_storage_moves",
	Method:  "GET",
}

var PostApiV4SnippetsRepositoryStorageMovesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/snippets/{id}/repository_storage_moves",
	Method:  "POST",
}

var GetApiV4SnippetsRepositoryStorageMovesByIdByRepositoryStorageMoveId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/snippets/{id}/repository_storage_moves/{repository_storage_move_id}",
	Method:  "GET",
}

var GetApiV4SnippetsUserAgentDetailById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/snippets/{id}/user_agent_detail",
	Method:  "GET",
}

var PutApiV4SuggestionsBatchApply EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/suggestions/batch_apply",
	Method:  "PUT",
}

var PutApiV4SuggestionsApplyById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/suggestions/{id}/apply",
	Method:  "PUT",
}

var GetApiV4Topics EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/topics",
	Method:  "GET",
}

var PostApiV4Topics EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/topics",
	Method:  "POST",
}

var PostApiV4TopicsMerge EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/topics/merge",
	Method:  "POST",
}

var DeleteApiV4TopicsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/topics/{id}",
	Method:  "DELETE",
}

var GetApiV4TopicsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/topics/{id}",
	Method:  "GET",
}

var PutApiV4TopicsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/topics/{id}",
	Method:  "PUT",
}

var PostApiV4UsageDataIncrementCounter EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/usage_data/increment_counter",
	Method:  "POST",
}

var PostApiV4UsageDataIncrementUniqueUsers EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/usage_data/increment_unique_users",
	Method:  "POST",
}

var GetApiV4UsageDataMetricDefinitions EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/usage_data/metric_definitions",
	Method:  "GET",
}

var GetApiV4UsageDataNonSqlMetrics EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/usage_data/non_sql_metrics",
	Method:  "GET",
}

var GetApiV4UsageDataQueries EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/usage_data/queries",
	Method:  "GET",
}

var GetApiV4UsageDataServicePing EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/usage_data/service_ping",
	Method:  "GET",
}

var PostApiV4UsageDataTrackEvent EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/usage_data/track_event",
	Method:  "POST",
}

var PostApiV4UsageDataTrackEvents EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/usage_data/track_events",
	Method:  "POST",
}

var GetApiV4UserActivities EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/activities",
	Method:  "GET",
}

var PutApiV4UserAvatar EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/avatar",
	Method:  "PUT",
}

var GetApiV4UserEmails EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/emails",
	Method:  "GET",
}

var PostApiV4UserEmails EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/emails",
	Method:  "POST",
}

var DeleteApiV4UserEmailsByEmailId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/emails/{email_id}",
	Method:  "DELETE",
}

var GetApiV4UserEmailsByEmailId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/emails/{email_id}",
	Method:  "GET",
}

var GetApiV4UserGpgKeys EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/gpg_keys",
	Method:  "GET",
}

var PostApiV4UserGpgKeys EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/gpg_keys",
	Method:  "POST",
}

var DeleteApiV4UserGpgKeysByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/gpg_keys/{key_id}",
	Method:  "DELETE",
}

var GetApiV4UserGpgKeysByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/gpg_keys/{key_id}",
	Method:  "GET",
}

var PostApiV4UserGpgKeysRevokeByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/gpg_keys/{key_id}/revoke",
	Method:  "POST",
}

var GetApiV4UserKeys EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/keys",
	Method:  "GET",
}

var PostApiV4UserKeys EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/keys",
	Method:  "POST",
}

var DeleteApiV4UserKeysByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/keys/{key_id}",
	Method:  "DELETE",
}

var GetApiV4UserKeysByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/keys/{key_id}",
	Method:  "GET",
}

var PostApiV4UserPersonalAccessTokens EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/personal_access_tokens",
	Method:  "POST",
}

var GetApiV4UserPreferences EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/preferences",
	Method:  "GET",
}

var PutApiV4UserPreferences EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/preferences",
	Method:  "PUT",
}

var PostApiV4UserRunners EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/runners",
	Method:  "POST",
}

var GetApiV4UserStatus EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/status",
	Method:  "GET",
}

var PatchApiV4UserStatus EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/status",
	Method:  "PATCH",
}

var PutApiV4UserStatus EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/status",
	Method:  "PUT",
}

var PutApiV4UserCreditCardValidationByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user/{user_id}/credit_card_validation",
	Method:  "PUT",
}

var GetApiV4UserCounts EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/user_counts",
	Method:  "GET",
}

var GetApiV4Users EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users",
	Method:  "GET",
}

var PostApiV4Users EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users",
	Method:  "POST",
}

var DeleteApiV4UsersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}",
	Method:  "DELETE",
}

var GetApiV4UsersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}",
	Method:  "GET",
}

var PutApiV4UsersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}",
	Method:  "PUT",
}

var PostApiV4UsersActivateById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/activate",
	Method:  "POST",
}

var PostApiV4UsersApproveById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/approve",
	Method:  "POST",
}

var GetApiV4UsersAssociationsCountById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/associations_count",
	Method:  "GET",
}

var PostApiV4UsersBanById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/ban",
	Method:  "POST",
}

var PostApiV4UsersBlockById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/block",
	Method:  "POST",
}

var GetApiV4UsersCustomAttributesById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/custom_attributes",
	Method:  "GET",
}

var DeleteApiV4UsersCustomAttributesByIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/custom_attributes/{key}",
	Method:  "DELETE",
}

var GetApiV4UsersCustomAttributesByIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/custom_attributes/{key}",
	Method:  "GET",
}

var PutApiV4UsersCustomAttributesByIdByKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/custom_attributes/{key}",
	Method:  "PUT",
}

var PostApiV4UsersDeactivateById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/deactivate",
	Method:  "POST",
}

var PatchApiV4UsersDisableTwoFactorById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/disable_two_factor",
	Method:  "PATCH",
}

var GetApiV4UsersEmailsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/emails",
	Method:  "GET",
}

var PostApiV4UsersEmailsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/emails",
	Method:  "POST",
}

var DeleteApiV4UsersEmailsByIdByEmailId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/emails/{email_id}",
	Method:  "DELETE",
}

var GetApiV4UsersEventsById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/events",
	Method:  "GET",
}

var PostApiV4UsersFollowById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/follow",
	Method:  "POST",
}

var GetApiV4UsersFollowersById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/followers",
	Method:  "GET",
}

var GetApiV4UsersFollowingById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/following",
	Method:  "GET",
}

var GetApiV4UsersGpgKeysById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/gpg_keys",
	Method:  "GET",
}

var PostApiV4UsersGpgKeysById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/gpg_keys",
	Method:  "POST",
}

var DeleteApiV4UsersGpgKeysByIdByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/gpg_keys/{key_id}",
	Method:  "DELETE",
}

var GetApiV4UsersGpgKeysByIdByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/gpg_keys/{key_id}",
	Method:  "GET",
}

var PostApiV4UsersGpgKeysRevokeByIdByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/gpg_keys/{key_id}/revoke",
	Method:  "POST",
}

var DeleteApiV4UsersIdentitiesByIdByProvider EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/identities/{provider}",
	Method:  "DELETE",
}

var DeleteApiV4UsersKeysByIdByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/keys/{key_id}",
	Method:  "DELETE",
}

var GetApiV4UsersKeysByIdByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/keys/{key_id}",
	Method:  "GET",
}

var PostApiV4UsersRejectById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/reject",
	Method:  "POST",
}

var PostApiV4UsersUnbanById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/unban",
	Method:  "POST",
}

var PostApiV4UsersUnblockById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/unblock",
	Method:  "POST",
}

var PostApiV4UsersUnfollowById EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{id}/unfollow",
	Method:  "POST",
}

var GetApiV4UsersContributedProjectsByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{user_id}/contributed_projects",
	Method:  "GET",
}

var GetApiV4UsersImpersonationTokensByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{user_id}/impersonation_tokens",
	Method:  "GET",
}

var PostApiV4UsersImpersonationTokensByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{user_id}/impersonation_tokens",
	Method:  "POST",
}

var DeleteApiV4UsersImpersonationTokensByUserIdByImpersonationTokenId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{user_id}/impersonation_tokens/{impersonation_token_id}",
	Method:  "DELETE",
}

var GetApiV4UsersImpersonationTokensByUserIdByImpersonationTokenId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{user_id}/impersonation_tokens/{impersonation_token_id}",
	Method:  "GET",
}

var GetApiV4UsersKeysByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{user_id}/keys",
	Method:  "GET",
}

var PostApiV4UsersKeysByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{user_id}/keys",
	Method:  "POST",
}

var GetApiV4UsersMembershipsByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{user_id}/memberships",
	Method:  "GET",
}

var PostApiV4UsersPersonalAccessTokensByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{user_id}/personal_access_tokens",
	Method:  "POST",
}

var GetApiV4UsersProjectDeployKeysByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{user_id}/project_deploy_keys",
	Method:  "GET",
}

var GetApiV4UsersProjectsByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{user_id}/projects",
	Method:  "GET",
}

var GetApiV4UsersStarredProjectsByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{user_id}/starred_projects",
	Method:  "GET",
}

var GetApiV4UsersStatusByUserId EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/users/{user_id}/status",
	Method:  "GET",
}

var GetApiV4Version EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/version",
	Method:  "GET",
}

var GetApiV4WebCommitsPublicKey EndpointPattern = EndpointPattern{
	Pattern: "/api/v4/web_commits/public_key",
	Method:  "GET",
}
